package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
)

type IJobService interface {
}

type IJobStageService interface {
	ProcessStages(ctx context.Context, jobId uuid.UUID) error
}

type IDownloadService interface {
	GetByHash(ctx context.Context, hash string) (*model.Download, error)
	SaveOne(ctx context.Context, download *model.Download) error
}

type ILockService interface {
	TryToLock(ctx context.Context, types, value string) bool
	TryToUnlock(ctx context.Context, types, value string) bool
}
