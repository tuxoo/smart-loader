package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
	"time"
)

type JobService struct {
	repository repository.IJobRepository
}

func NewJobService(repository repository.IJobRepository) *JobService {
	return &JobService{
		repository: repository,
	}
}

func (s *JobService) Create(ctx context.Context, uris []string) (model.JobStatusDto, error) {
	job := model.Job{
		Id:        uuid.New(),
		Name:      "TEST",
		Size:      len(uris),
		Status:    model.NEW,
		CreatedAt: time.Now(),
	}

	if err := s.repository.Save(ctx, job); err != nil {
		return model.JobStatusDto{}, err
	}

	return model.JobStatusDto{
		Id:        job.Id,
		Status:    job.Status,
		CreatedAt: job.CreatedAt,
	}, nil
}
