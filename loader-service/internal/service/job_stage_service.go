package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/downloader"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/hasher"
	"strconv"
	"time"
)

type JobStageService struct {
	repository      repository.IJobStageRepository
	downloadService IDownloadService
	minioService    IMinioService
	lockService     ILockService
	downloader      downloader.Downloader
	hasher          hasher.Hasher
}

func NewJobStageService(
	repository repository.IJobStageRepository,
	downloadService IDownloadService,
	minioService IMinioService,
	lockService ILockService,
	downloader downloader.Downloader,
	hasher hasher.Hasher,
) *JobStageService {
	return &JobStageService{
		repository:      repository,
		downloadService: downloadService,
		lockService:     lockService,
		minioService:    minioService,
		downloader:      downloader,
		hasher:          hasher,
	}
}

func (s *JobStageService) ProcessStages(ctx context.Context, jobId uuid.UUID) error {
	stages, err := s.repository.FindAllByJobId(ctx, jobId)
	if err != nil {
		return err
	}

	for _, stage := range stages {
		if ok := s.lockService.TryToLock(ctx, model.JOB_STAGE_LOCK, strconv.Itoa(stage.Id)); ok {
			if err = s.processingStage(ctx, &stage); err != nil {
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
	urls := stage.Urls
	for _, url := range urls {
		bytes, err := s.downloader.Download(url)
		if err != nil {
			return err
		}

		hash := s.hasher.HashBytes(bytes)
		download, err := s.downloadService.GetByHash(ctx, hash)
		if err != nil {
			return err
		}

		if download == nil {
			download = &model.Download{
				Id:           uuid.New(),
				Hash:         hash,
				DownloadedAt: time.Now(),
			}

			if err = s.downloadService.SaveOne(ctx, download); err != nil {
				return err
			}
		}

		// TODO: save to Minio

		fmt.Println(url, len(bytes))
	}

	return nil
}
