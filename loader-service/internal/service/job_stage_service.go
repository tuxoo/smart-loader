package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model/const"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/downloader"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/hasher"
	"io"
	"strconv"
	"sync"
	"time"
)

type JobStageService struct {
	repository              repository.IJobStageRepository
	downloadService         IDownloadService
	jobService              IJobService
	jobStageDownloadService IJobStageDownloadService
	minioService            IMinioService
	lockService             ILockService
	downloader              downloader.Downloader
	hasher                  hasher.Hasher
}

func NewJobStageService(
	repository repository.IJobStageRepository,
	downloadService IDownloadService,
	jobService IJobService,
	jobStageDownloadService IJobStageDownloadService,
	minioService IMinioService,
	lockService ILockService,
	downloader downloader.Downloader,
	hasher hasher.Hasher,
) *JobStageService {
	return &JobStageService{
		repository:              repository,
		downloadService:         downloadService,
		jobService:              jobService,
		jobStageDownloadService: jobStageDownloadService,
		lockService:             lockService,
		minioService:            minioService,
		downloader:              downloader,
		hasher:                  hasher,
	}
}

func (s *JobStageService) ProcessStages(ctx context.Context, jobId uuid.UUID) error {
	if err := s.jobService.UpdateStatus(ctx, jobId, _const.PENDING_STATUS); err != nil {
		return err
	}

	stages, err := s.repository.FindAllByJobId(ctx, jobId)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, stage := range stages {
		wg.Add(1)

		go func(stage model.BriefJobStage) {
			defer wg.Done()

			var status string
			if ok := s.lockService.TryToLock(ctx, _const.JOB_STAGE_LOCK, strconv.Itoa(stage.Id)); ok {
				if err = s.processingStage(ctx, &stage); err != nil {
					status = _const.FAILED_STATUS
				} else {
					status = _const.EXECUTED_STATUS
				}

				if err = s.repository.UpdateStatus(ctx, stage.Id, status); err != nil {
					// TODO:
				}

				s.lockService.TryToUnlock(ctx, _const.JOB_STAGE_LOCK, strconv.Itoa(stage.Id))
			}
		}(stage)
	}

	wg.Wait()

	stages, err = s.repository.FindAllByJobId(ctx, jobId)
	if err != nil {
		return err
	}

	return s.jobService.UpdateStatusByStages(ctx, jobId, stages)
}

func (s *JobStageService) processingStage(ctx context.Context, stage *model.BriefJobStage) error {
	for _, url := range stage.Urls {
		if err := s.repository.UpdateStatus(ctx, stage.Id, _const.PENDING_STATUS); err != nil {
			return err
		}

		reader, err := s.downloader.Download(url)
		if err != nil {
			return err
		}

		content, err := io.ReadAll(reader)
		if err != nil {
			return err
		}

		if err = reader.Close(); err != nil {
			return err
		}

		hash := s.hasher.HashBytes(content)
		download, err := s.downloadService.GetByHash(ctx, hash)
		if err != nil {
			return err
		}

		if download != nil {
			if err = s.jobStageDownloadService.Save(ctx, stage.Id, download.Id); err != nil {
				return err
			}
			continue
		}

		download = &model.Download{
			Id:           uuid.New(),
			Hash:         hash,
			Size:         len(content),
			DownloadedAt: time.Now(),
		}

		if err = s.minioService.Put(ctx, content, download); err != nil {
			return err
		}

		tx, err := s.repository.CreateTransaction(ctx)

		if err = s.downloadService.SaveOne(ctx, tx, download); err != nil {
			if err = tx.Rollback(ctx); err != nil {
				return err
			}
		}

		if err = s.jobStageDownloadService.SaveInTransaction(ctx, tx, stage.Id, download.Id); err != nil {
			if err = tx.Rollback(ctx); err != nil {
				return err
			}
		}

		if err = tx.Commit(ctx); err != nil {
			return err
		}
	}

	return nil
}
