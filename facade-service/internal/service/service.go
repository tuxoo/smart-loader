package service

import (
	"context"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
)

type IJobService interface {
	Create(ctx context.Context, uris []string) (model.JobStatusDto, error)
}

type Services struct {
	JobService IJobService
}

type ServicesDeps struct {
	Repositories *repository.Repositories
}

func NewServices(deps ServicesDeps) *Services {
	jobService := NewJobService(deps.Repositories.JobRepository)

	return &Services{
		JobService: jobService,
	}
}
