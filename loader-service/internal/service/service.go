package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
)

type IJobService interface {
}

type IJobStageService interface {
	ProcessStages(ctx context.Context, jobId uuid.UUID) error
}

type IDownloadService interface {
	GetByHash(ctx context.Context, hash string) (*model.Download, error)
	SaveOne(ctx context.Context, tx pgx.Tx, download *model.Download) error
}

type IJobStageDownloadService interface {
	Save(ctx context.Context, jobStageId int, downloadId uuid.UUID) error
	SaveInTransaction(ctx context.Context, tx pgx.Tx, jobStageId int, downloadId uuid.UUID) error
}

type IMinioService interface {
	Put(ctx context.Context, content []byte, download *model.Download) error
	Get() error
}

type ILockService interface {
	TryToLock(ctx context.Context, types, value string) bool
	TryToUnlock(ctx context.Context, types, value string) bool
}
