package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

type IUserService interface {
	SignIn(ctx context.Context, dto model.SignInDTO) (string, error)
}

type IJobService interface {
	Create(ctx context.Context, userId int, urls []string) (job model.Job, err error)
	GetAll(ctx context.Context, userId int) ([]model.Job, error)
}

type IJobStageService interface {
	Create(ctx context.Context, tx pgx.Tx, jobId uuid.UUID, urls []string) error
	GetAllByJobId(ctx context.Context, jobId uuid.UUID) ([]int, error)
}
