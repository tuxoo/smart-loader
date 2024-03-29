package setup

import (
	"github.com/tuxoo/smart-loader/facade-service/internal/client"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/facade-service/internal/service"
	"github.com/tuxoo/smart-loader/facade-service/internal/util/hasher"
	"github.com/tuxoo/smart-loader/facade-service/internal/util/token-manager"
)

func provideUserService(
	repository repository.IUserRepository,
	hasher hasher.Hasher,
	tokenManager token_manager.TokenManager,
	tokenService service.ITokenService,
) service.IUserService {
	return service.NewUserService(repository, hasher, tokenManager, tokenService)
}

func provideTokenService(
	cfg *config.AppConfig,
	repository repository.ITokenRepository,
) service.ITokenService {
	return service.NewTokenService(cfg, repository)
}

func provideJobService(
	repository repository.IJobRepository,
	jobStageService service.IJobStageService,
	natsClient *client.NatsClient,
) service.IJobService {
	return service.NewJobService(repository, jobStageService, natsClient)
}

func provideJobStageService(
	cfg *config.AppConfig,
	repository repository.IJobStageRepository,
) service.IJobStageService {
	return service.NewJobStageService(cfg, repository)
}

func provideDownloadService(
	repository repository.IDownloadRepository,
	minioService service.IMinioService,
) service.IDownloadService {
	return service.NewDownloadService(repository, minioService)
}

func provideMinioService(client *client.MinioClient) service.IMinioService {
	return service.NewMinioService(client)
}
