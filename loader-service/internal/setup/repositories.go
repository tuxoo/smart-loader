package setup

import "github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"

func provideJobRepository(db *repository.PostgresDB) repository.IJobRepository {
	return repository.NewJobRepository(db)
}

func provideJobStageRepository(db *repository.PostgresDB) repository.IJobStageRepository {
	return repository.NewJobStageRepository(db)
}

func provideLockeRepository(db *repository.PostgresDB) repository.ILockRepository {
	return repository.NewLockRepository(db)
}
