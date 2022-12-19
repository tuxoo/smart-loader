package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
)

type IJobRepository interface {
}

type IJobStageRepository interface {
	FindAllByJobId(ctx context.Context, jobId uuid.UUID) ([]model.BriefJobStage, error)
}

type IDownloadRepository interface {
	FindByHash(ctx context.Context, hash string) (*model.Download, error)
	SaveOne(ctx context.Context, download *model.Download) error
}

type ILockRepository interface {
	Lock(ctx context.Context, types, value string) error
	Unlock(ctx context.Context, types, value string) error
}
