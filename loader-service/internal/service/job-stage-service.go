package service

import (
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
