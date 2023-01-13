package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

type IUserRepository interface {
	FindByCredentials(ctx context.Context, email, password string) (model.User, error)
}

type IJobRepository interface {
	CreateTransaction(ctx context.Context) (pgx.Tx, error)
	FindAll(ctx context.Context, userId int) ([]model.Job, error)
	SaveInTransaction(ctx context.Context, tx pgx.Tx, job model.Job) error
}

type IJobStageRepository interface {
	SaveInTransaction(ctx context.Context, tx pgx.Tx, jobStage model.JobStage) error
	FindAllByJobId(ctx context.Context, jobId uuid.UUID) ([]int, error)
}
