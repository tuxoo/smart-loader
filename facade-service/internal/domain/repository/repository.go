package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

const (
	jobTable              = "job"
	jobStageTable         = "job_stage"
	jobStageDownloadTable = "job_stage_download"
	downloadTable         = "download"
	userTable             = "\"user\""
)

type (
	IUserRepository interface {
		FindByCredentials(ctx context.Context, email, password string) (model.User, error)
		UpdateLastVisit(ctx context.Context, id int) (err error)
		FindById(ctx context.Context, id int) (model.User, error)
	}

	IJobRepository interface {
		CreateTransaction(ctx context.Context) (pgx.Tx, error)
		FindAll(ctx context.Context, userId int) ([]model.Job, error)
		SaveInTransaction(ctx context.Context, tx pgx.Tx, job model.Job) error
	}

	IJobStageRepository interface {
		SaveInTransaction(ctx context.Context, tx pgx.Tx, jobStage model.JobStage) error
		FindAllByJobId(ctx context.Context, jobId uuid.UUID) ([]int, error)
	}

	IDownloadRepository interface {
		FindAllByJobId(ctx context.Context, jobId uuid.UUID, userId int) ([]model.Download, error)
	}
)
