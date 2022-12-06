package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/loader-service/internal/model"
)

const (
	userTable     = "\"user\""
	jobTable      = "job"
	jobStageTable = "job_stage"
	downloadTable = "download"
)

type IJobRepository interface {
	CreateTransaction(ctx context.Context) (pgx.Tx, error)
	Save(ctx context.Context, tx pgx.Tx, job model.Job) error
}

type IJobStageRepository interface {
	Save(ctx context.Context, tx pgx.Tx, jobStage model.JobStage) error
}

type Repositories struct {
	JobRepository      IJobRepository
	JobStageRepository IJobStageRepository
}

func NewRepositories(db *PostgresDB) *Repositories {
	return &Repositories{
		JobRepository:      NewJobRepository(db.pool),
		JobStageRepository: NewJobStageRepository(db.pool),
	}
}
