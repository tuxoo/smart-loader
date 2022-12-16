package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	model2 "github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
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

func (s *JobStageService) Create(ctx context.Context, tx pgx.Tx, jobId uuid.UUID, uris []string) error {
	urisPartitions := partitionUris(uris, s.cfg.UriPartitionSize)

	for _, partition := range urisPartitions {
		jobStage := model2.JobStage{
			Id:     uuid.New(),
			Size:   len(partition),
			Uris:   partition,
			Status: model2.NEW,
			JobId:  jobId,
		}

		if err := s.repository.Save(ctx, tx, jobStage); err != nil {
			return err
		}
	}

	return nil
}

func partitionUris(uris []string, partitionSize int) (partitions [][]string) {
	for {
		if len(uris) == 0 {
			break
		}

		if len(uris) < partitionSize {
			partitionSize = len(uris)
		}

		partitions = append(partitions, uris[0:partitionSize])
		uris = uris[partitionSize:]
	}

	return
}
