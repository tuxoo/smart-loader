package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
)

type IJobService interface {
	Create(ctx context.Context, uris []string) (model.JobStatusDto, error)
}

type IJobStageService interface {
	Create(ctx context.Context, jobId uuid.UUID, uris []string) error
}

type Services struct {
	JobService      IJobService
	JobStageService IJobStageService
}

type ServicesDeps struct {
	Repositories *repository.Repositories
}

func NewServices(deps ServicesDeps) *Services {
	jobStageService := NewJobStageService(deps.Repositories.JobStageRepository)

	return &Services{
		JobService:      NewJobService(deps.Repositories.JobRepository, jobStageService),
		JobStageService: jobStageService,
	}
}
