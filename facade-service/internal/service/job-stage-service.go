package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
)

type JobStageService struct {
	repository repository.IJobStageRepository
}

func NewJobStageService(repository repository.IJobStageRepository) *JobStageService {
	return &JobStageService{
		repository: repository,
	}
}

func (s *JobStageService) Create(ctx context.Context, tx pgx.Tx, jobId uuid.UUID, uris []string) error {
	urisPartitions := partitionUris(uris, 2)

	for _, partition := range urisPartitions {
		jobStage := model.JobStage{
			Id:     uuid.New(),
			Size:   len(partition),
			Uris:   partition,
			Status: model.NEW,
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
