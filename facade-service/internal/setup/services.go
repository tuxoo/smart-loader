package setup

import (
	"github.com/tuxoo/smart-loader/facade-service/internal/client"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/facade-service/internal/service"
	"github.com/tuxoo/smart-loader/facade-service/internal/util/hasher"
	"github.com/tuxoo/smart-loader/facade-service/internal/util/token-manager"
)

func provideUserService(repository repository.IUserRepository, hasher hasher.Hasher, tokenManager token_manager.TokenManager) service.IUserService {
	return service.NewUserService(repository, hasher, tokenManager)
}

func provideJobService(repository repository.IJobRepository, jobStageService service.IJobStageService, natsClient *client.NatsClient) service.IJobService {
	return service.NewJobService(repository, jobStageService, natsClient)
}

func provideJobStageService(cfg *config.AppConfig, repository repository.IJobStageRepository) service.IJobStageService {
	return service.NewJobStageService(cfg, repository)
}
