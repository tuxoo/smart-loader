package setup

import "github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"

func provideJobRepository(db *repository.PostgresDB) repository.IJobRepository {
	return repository.NewJobRepository(db)
}

func provideJobStageRepository(db *repository.PostgresDB) repository.IJobStageRepository {
	return repository.NewJobStageRepository(db)
}

func provideJobStageDownloadRepository(db *repository.PostgresDB) repository.IJobStageDownloadRepository {
	return repository.NewJobStageDownloadRepository(db)
}

func provideDownloadRepository(db *repository.PostgresDB) repository.IDownloadRepository {
	return repository.NewDownloadRepository(db)
}

func provideLockRepository(db *repository.PostgresDB) repository.ILockRepository {
	return repository.NewLockRepository(db)
}
