package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
)

type IJobService interface {
	Create(ctx context.Context, uris []string) (model.JobStatusDto, error)
}

type IJobStageService interface {
	Create(ctx context.Context, tx pgx.Tx, jobId uuid.UUID, uris []string) error
}

type Services struct {
	JobService      IJobService
	JobStageService IJobStageService
}

func NewServices(repositories *repository.Repositories) *Services {
	jobStageService := NewJobStageService(repositories.JobStageRepository)

	return &Services{
		JobService:      NewJobService(repositories.JobRepository, jobStageService),
		JobStageService: jobStageService,
	}
}
