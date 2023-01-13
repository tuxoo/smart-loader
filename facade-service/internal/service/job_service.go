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

const NEW_JOB = "job.new"

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

// TODO: regexp for URLs
func (s *JobService) Create(ctx context.Context, userId int, urls []string) (job model.Job, err error) {
	tx, err := s.repository.CreateTransaction(ctx)
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	job = model.Job{
		Id:        uuid.New(),
		Size:      len(urls),
		Status:    model.NEW,
		CreatedAt: time.Now(),
		UserId:    userId,
	}

	err = s.repository.SaveInTransaction(ctx, tx, job)
	if err != nil {
		return job, err
	}

	if err = s.jobStageService.Create(ctx, tx, job.Id, urls); err != nil {
		return job, err
	}

	if err = tx.Commit(ctx); err != nil {
		return job, err
	}

	err = s.natsClient.Conn.Publish(NEW_JOB, []byte(job.Id.String()))
	if err != nil {
		return job, err
	}

	return job, nil
}

func (s *JobService) GetAll(ctx context.Context, userId int) ([]model.Job, error) {
	return s.repository.FindAll(ctx, userId)
}
