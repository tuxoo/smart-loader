package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	model2 "github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

const (
	userTable     = "\"user\""
	jobTable      = "job"
	jobStageTable = "job_stage"
	downloadTable = "download"
)

type IUserRepository interface {
	FindByCredentials(ctx context.Context, email, password string) (*model2.User, error)
}

type IJobRepository interface {
	CreateTransaction(ctx context.Context) (pgx.Tx, error)
	Save(ctx context.Context, tx pgx.Tx, job model2.Job) error
}

type IJobStageRepository interface {
	Save(ctx context.Context, tx pgx.Tx, jobStage model2.JobStage) error
}

type Repositories struct {
	UserRepository     IUserRepository
	JobRepository      IJobRepository
	JobStageRepository IJobStageRepository
}
