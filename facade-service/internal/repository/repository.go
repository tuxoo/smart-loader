package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
)

const (
	userTable     = "\"user\""
	jobTable      = "job"
	jobStageTable = "job_stage"
	downloadTable = "download"
)

type IJobRepository interface {
	Save(ctx context.Context, job model.Job) error
}

type Repositories struct {
	JobRepository IJobRepository
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		JobRepository: NewJobRepository(db),
	}
}
