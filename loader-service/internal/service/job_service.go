package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
	_const "github.com/tuxoo/smart-loader/loader-service/internal/domain/model/const"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
)

type JobService struct {
	repository repository.IJobRepository
}

func NewJobService(repository repository.IJobRepository) *JobService {
	return &JobService{
		repository: repository,
	}
}

func (s *JobService) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return s.repository.UpdateStatus(ctx, id, status)
}

func (s *JobService) UpdateStatusByStages(ctx context.Context, id uuid.UUID, stages []model.BriefJobStage) error {
	var status string
	for _, stage := range stages {
		switch stage.Status {
		case _const.NEW_STATUS:
			status = _const.PARTLY_EXECUTED_STATUS
			break
		case _const.FAILED_STATUS:
			status = _const.PARTLY_EXECUTED_STATUS
			break
		case _const.PENDING_STATUS:
			status = _const.PARTLY_EXECUTED_STATUS
			break
		default:
			status = _const.EXECUTED_STATUS
		}
	}

	return s.repository.UpdateStatus(ctx, id, status)
}
