package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/minio/minio-go/v7"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

type (
	IUserService interface {
		SignIn(ctx context.Context, dto model.SignInDTO) (string, error)
	}

	IJobService interface {
		Create(ctx context.Context, userId int, urls []string) (job model.Job, err error)
		GetAll(ctx context.Context, userId int) ([]model.Job, error)
	}

	IJobStageService interface {
		Create(ctx context.Context, tx pgx.Tx, jobId uuid.UUID, urls []string) error
		GetAllByJobId(ctx context.Context, jobId uuid.UUID) ([]int, error)
	}

	IDownloadService interface {
		GetDownloadZip(ctx context.Context, jobId uuid.UUID, userId int) ([]byte, error)
	}

	IMinioService interface {
		Get(ctx context.Context, downloadId uuid.UUID) (*minio.Object, error)
	}
)
