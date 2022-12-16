package service

import (
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
)

type JobService struct {
	repository repository.IJobRepository
}

func NewJobService(repository repository.IJobRepository) *JobService {
	return &JobService{
		repository: repository,
	}
}
