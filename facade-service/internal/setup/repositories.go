package setup

import "github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"

func provideUserRepository(db *repository.PostgresDB) repository.IUserRepository {
	return repository.NewUserRepository(db)
}

func provideTokenRepository(db *repository.PostgresDB) repository.ITokenRepository {
	return repository.NewTokenRepository(db)
}

func provideJobRepository(db *repository.PostgresDB) repository.IJobRepository {
	return repository.NewJobRepository(db)
}

func provideJobStageRepository(db *repository.PostgresDB) repository.IJobStageRepository {
	return repository.NewJobStageRepository(db)
}

func provideDownloadRepository(db *repository.PostgresDB) repository.IDownloadRepository {
	return repository.NewDownloadRepository(db)
}
