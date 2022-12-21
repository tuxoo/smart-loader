package service

import (
	"context"
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

func (s *JobService) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return s.repository.UpdateStatus(ctx, id, status)
}
