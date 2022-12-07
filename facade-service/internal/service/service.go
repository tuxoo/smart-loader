package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
	"github.com/tuxoo/smart-loader/facade-service/internal/util"
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

func NewServices(
	cfg *config.AppConfig,
	repositories *repository.Repositories,
	hasher *util.Hasher,
) *Services {
	jobStageService := NewJobStageService(cfg, repositories.JobStageRepository)

	return &Services{
		UserService:     NewUserService(repositories.UserRepository, hasher),
		JobService:      NewJobService(repositories.JobRepository, jobStageService),
		JobStageService: jobStageService,
	}
}
