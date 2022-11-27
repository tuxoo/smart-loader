package repository

import (
	"context"
	"github.com/google/uuid"
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
	Save(ctx context.Context, job model.Job) (uuid.UUID, error)
}

type IJobStageRepository interface {
	Save(ctx context.Context, jobStage model.JobStage) error
}

type Repositories struct {
	JobRepository      IJobRepository
	JobStageRepository IJobStageRepository
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		JobRepository:      NewJobRepository(db),
		JobStageRepository: NewJobStageRepository(db),
	}
}
