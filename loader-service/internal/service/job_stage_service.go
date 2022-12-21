package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/downloader"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/hasher"
	"strconv"
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
	if err := s.jobService.UpdateStatus(ctx, jobId, model.PENDING_STATUS); err != nil {
		return err
	}

	stages, err := s.repository.FindAllByJobId(ctx, jobId)
	if err != nil {
		return err
	}

	for _, stage := range stages {
		if ok := s.lockService.TryToLock(ctx, model.JOB_STAGE_LOCK, strconv.Itoa(stage.Id)); ok {
			if err = s.repository.UpdateStatus(ctx, stage.Id, model.PENDING_STATUS); err != nil {
				return err
			}

			if err = s.processingStage(ctx, &stage); err != nil {
				if err = s.repository.UpdateStatus(ctx, stage.Id, model.FAILED_STATUS); err != nil {
					return err
				}
				return err
			}

			if err = s.repository.UpdateStatus(ctx, stage.Id, model.EXECUTED_STATUS); err != nil {
				return err
			}

			s.lockService.TryToUnlock(ctx, model.JOB_STAGE_LOCK, strconv.Itoa(stage.Id))
		} else {
			continue
		}
	}

	return nil
}

func (s *JobStageService) processingStage(ctx context.Context, stage *model.BriefJobStage) error {
	for _, url := range stage.Urls {
		content, err := s.downloader.Download(url)
		if err != nil {
			return err
		}

		hash := s.hasher.HashBytes(content)
		download, err := s.downloadService.GetByHash(ctx, hash)
		if err != nil {
			continue
		}

		if download != nil {
			_ = s.jobStageDownloadService.Save(ctx, stage.Id, download.Id)
			continue
		}

		download = &model.Download{
			Id:           uuid.New(),
			Hash:         hash,
			Size:         len(content),
			DownloadedAt: time.Now(),
		}

		if err = s.minioService.Put(ctx, content, download); err != nil {
			continue
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
