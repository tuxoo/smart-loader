package service

import (
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
)

type JobStageService struct {
	repository repository.IJobStageRepository
}

func NewJobStageService(repository repository.IJobStageRepository) *JobStageService {
	return &JobStageService{
		repository: repository,
	}
}

func (s *JobStageService) ProcessStage(jobId uuid.UUID) error {

	return nil
}
