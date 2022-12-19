package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
)

type JobStageService struct {
	repository  repository.IJobStageRepository
	lockService ILockService
}

func NewJobStageService(repository repository.IJobStageRepository, lockService ILockService) *JobStageService {
	return &JobStageService{
		repository:  repository,
		lockService: lockService,
	}
}

func (s *JobStageService) ProcessStage(ctx context.Context, jobId uuid.UUID) error {
	s.lockService.TryToLock(ctx, model.JOB_STAGE_LOCK, jobId.String())

	s.lockService.TryToUnlock(ctx, model.JOB_STAGE_LOCK, jobId.String())
	return nil
}
