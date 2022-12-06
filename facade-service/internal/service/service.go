package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
)

type IUserService interface {
	SignIn(ctx context.Context, dto model.SignInDTO) (string, error)
}

type IJobService interface {
	Create(ctx context.Context, uris []string) (model.JobStatusDto, error)
}

type IJobStageService interface {
	Create(ctx context.Context, tx pgx.Tx, jobId uuid.UUID, uris []string) error
}

type Services struct {
	UserService     IUserService
	JobService      IJobService
	JobStageService IJobStageService
}

func NewServices(repositories *repository.Repositories, cfg *config.AppConfig) *Services {
	jobStageService := NewJobStageService(cfg, repositories.JobStageRepository)

	return &Services{
		UserService:     NewUserService(repositories.UserRepository),
		JobService:      NewJobService(repositories.JobRepository, jobStageService),
		JobStageService: jobStageService,
	}
}
