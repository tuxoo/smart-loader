package service

import (
	"fmt"
	"github.com/google/uuid"
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

func (s *JobService) ProcessJob(jobId uuid.UUID) error {

	fmt.Println(jobId)

	return nil
}
