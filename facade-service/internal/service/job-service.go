package service

import (
	"context"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
	"time"
)

type JobService struct {
	repository      repository.IJobRepository
	jobStageService IJobStageService
}

func NewJobService(repository repository.IJobRepository, jobStageService IJobStageService) *JobService {
	return &JobService{
		repository:      repository,
		jobStageService: jobStageService,
	}
}

// TODO: regexp for URIs
// TODO: add transactions
func (s *JobService) Create(ctx context.Context, uris []string) (model.JobStatusDto, error) {
	job := model.Job{
		Name:      "TEST",
		Size:      len(uris),
		Status:    model.NEW,
		CreatedAt: time.Now(),
	}

	jobId, err := s.repository.Save(ctx, job)
	if err != nil {
		return model.JobStatusDto{}, err
	}

	if err = s.jobStageService.Create(ctx, jobId, uris); err != nil {
		return model.JobStatusDto{}, err
	}

	return model.JobStatusDto{
		Id:        jobId,
		Status:    job.Status,
		CreatedAt: job.CreatedAt,
	}, nil
}
