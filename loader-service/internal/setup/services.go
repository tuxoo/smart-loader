package setup

import (
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/loader-service/internal/service"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/downloader"
	"github.com/tuxoo/smart-loader/loader-service/internal/util/hasher"
)

func provideJobService(repository repository.IJobRepository) service.IJobService {
	return service.NewJobService(repository)
}

func provideJobStageService(
	repository repository.IJobStageRepository,
	downloadService service.IDownloadService,
	lockService service.ILockService,
	downloader downloader.Downloader,
	hasher hasher.Hasher,
) service.IJobStageService {
	return service.NewJobStageService(repository, downloadService, lockService, downloader, hasher)
}

func provideDownloadService(repository repository.IDownloadRepository) service.IDownloadService {
	return service.NewDownloadService(repository)
}

func provideLockService(repository repository.ILockRepository) service.ILockService {
	return service.NewLockService(repository)
}
