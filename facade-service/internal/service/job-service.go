package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
	"time"
)

type JobService struct {
	repository      repository.IJobRepository
	jobStageService IJobStageService
}

func NewJobService(repository repository.IJobRepository, jobStageService IJobStageService) *JobService {
	return &JobService{
		repository:      repository,
		jobStageService: jobStageService,
	}
}

// TODO: regexp for URIs
func (s *JobService) Create(ctx context.Context, uris []string) (model.JobStatusDto, error) {
	tx, err := s.repository.CreateTransaction(ctx)
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	job := model.Job{
		Id:        uuid.New(),
		Name:      "TEST",
		Size:      len(uris),
		Status:    model.NEW,
		CreatedAt: time.Now(),
	}

	err = s.repository.Save(ctx, tx, job)
	if err != nil {
		return model.JobStatusDto{}, err
	}

	if err = s.jobStageService.Create(ctx, tx, job.Id, uris); err != nil {
		return model.JobStatusDto{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return model.JobStatusDto{}, err
	}

	return model.JobStatusDto{
		Id:        job.Id,
		Status:    job.Status,
		CreatedAt: job.CreatedAt,
	}, nil
}
