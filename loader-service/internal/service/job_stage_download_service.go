package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
)

type JobStageDownloadService struct {
	repository repository.IJobStageDownloadRepository
}

func NewJobStageDownloadService(repository repository.IJobStageDownloadRepository) *JobStageDownloadService {
	return &JobStageDownloadService{
		repository: repository,
	}
}

func (s *JobStageDownloadService) Save(ctx context.Context, jobStageId int, downloadId uuid.UUID) error {
	return s.repository.Save(ctx, jobStageId, downloadId)
}

func (s *JobStageDownloadService) SaveInTransaction(ctx context.Context, tx pgx.Tx, jobStageId int, downloadId uuid.UUID) error {
	return s.repository.SaveInTransaction(ctx, tx, jobStageId, downloadId)
}
