package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

const (
	userTable     = "\"user\""
	jobTable      = "job"
	jobStageTable = "job_stage"
	downloadTable = "download"
)

type IUserRepository interface {
	FindByCredentials(ctx context.Context, email, password string) (*model.User, error)
}

type IJobRepository interface {
	CreateTransaction(ctx context.Context) (pgx.Tx, error)
	Save(ctx context.Context, tx pgx.Tx, job model.Job) error
}

type IJobStageRepository interface {
	Save(ctx context.Context, tx pgx.Tx, jobStage model.JobStage) error
}
