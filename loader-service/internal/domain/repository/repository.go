package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
)

type (
	IJobRepository interface {
		UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
	}

	IJobStageRepository interface {
		CreateTransaction(ctx context.Context) (pgx.Tx, error)
		UpdateStatus(ctx context.Context, id int, status string) error
		UpdateStatusInTransaction(ctx context.Context, tx pgx.Tx, id int, status string) error
		FindAllByJobId(ctx context.Context, jobId uuid.UUID) ([]model.BriefJobStage, error)
	}

	IDownloadRepository interface {
		CreateTransaction(ctx context.Context) (pgx.Tx, error)
		FindByHash(ctx context.Context, hash string) (*model.Download, error)
		SaveOne(ctx context.Context, tx pgx.Tx, download *model.Download) error
	}

	IJobStageDownloadRepository interface {
		CreateTransaction(ctx context.Context) (pgx.Tx, error)
		Save(ctx context.Context, jobStageId int, downloadId uuid.UUID) error
		SaveInTransaction(ctx context.Context, tx pgx.Tx, jobStageId int, downloadId uuid.UUID) error
	}

	ILockRepository interface {
		Lock(ctx context.Context, types, value string) error
		Unlock(ctx context.Context, types, value string) error
	}
)
