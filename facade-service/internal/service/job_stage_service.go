package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"
)

type JobStageService struct {
	cfg        *config.AppConfig
	repository repository.IJobStageRepository
}

func NewJobStageService(cfg *config.AppConfig, repository repository.IJobStageRepository) *JobStageService {
	return &JobStageService{
		cfg:        cfg,
		repository: repository,
	}
}

func (s *JobStageService) Create(ctx context.Context, tx pgx.Tx, jobId uuid.UUID, urls []string) error {
	urlsPartitions := partitioningUrls(urls, s.cfg.UriPartitionSize)

	for _, partition := range urlsPartitions {
		jobStage := model.JobStage{
			Id:     uuid.New(),
			Size:   len(partition),
			Urls:   partition,
			Status: model.NEW,
			JobId:  jobId,
		}

		if err := s.repository.SaveInTransaction(ctx, tx, jobStage); err != nil {
			return err
		}
	}

	return nil
}

func (s *JobStageService) GetAllByJobId(ctx context.Context, jobId uuid.UUID) ([]int, error) {
	return s.repository.FindAllByJobId(ctx, jobId)
}

func partitioningUrls(urls []string, partitionSize int) (partitions [][]string) {
	for {
		if len(urls) == 0 {
			break
		}

		if len(urls) < partitionSize {
			partitionSize = len(urls)
		}

		partitions = append(partitions, urls[0:partitionSize])
		urls = urls[partitionSize:]
	}

	return
}
