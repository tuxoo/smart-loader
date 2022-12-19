package setup

import (
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/loader-service/internal/service"
)

func provideJobService(repository repository.IJobRepository) service.IJobService {
	return service.NewJobService(repository)
}

func provideJobStageService(repository repository.IJobStageRepository) service.IJobStageService {
	return service.NewJobStageService(repository)
}

func provideLockService(repository repository.ILockRepository) service.ILockService {
	return service.NewLockService(repository)
}
