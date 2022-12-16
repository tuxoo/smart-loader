package setup

import "github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"

func provideUserRepository(db *repository.PostgresDB) repository.IUserRepository {
	return repository.NewUserRepository(db)
}

func provideJobRepository(db *repository.PostgresDB) repository.IJobStageRepository {
	return repository.NewJobStageRepository(db)
}

func provideJobStageRepository(db *repository.PostgresDB) repository.IJobStageRepository {
	return repository.NewJobStageRepository(db)
}
