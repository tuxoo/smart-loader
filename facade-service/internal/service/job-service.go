package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/client"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"
	"time"
)

type JobService struct {
	repository      repository.IJobRepository
	jobStageService IJobStageService
	natsClient      *client.NatsClient
}

func NewJobService(repository repository.IJobRepository, jobStageService IJobStageService, natsClient *client.NatsClient) *JobService {
	return &JobService{
		repository:      repository,
		jobStageService: jobStageService,
		natsClient:      natsClient,
	}
}

// TODO: regexp for URIs
func (s *JobService) Create(ctx context.Context, userId int, uris []string) (*model.JobStatusDto, error) {
	tx, err := s.repository.CreateTransaction(ctx)
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	job := model.Job{
		Id:        uuid.New(),
		Size:      len(uris),
		Status:    model.NEW,
		CreatedAt: time.Now(),
		UserId:    userId,
	}

	err = s.repository.Save(ctx, tx, job)
	if err != nil {
		return nil, err
	}

	if err = s.jobStageService.Create(ctx, tx, job.Id, uris); err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	err = s.natsClient.Conn.Publish("foo", []byte("Hello World"))
	if err != nil {
		return nil, err
	}

	return &model.JobStatusDto{
		Id:        job.Id,
		Status:    job.Status,
		CreatedAt: job.CreatedAt,
	}, nil
}
