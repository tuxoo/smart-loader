package repository

const jobTable = "job"

type JobRepository struct {
	db *PostgresDB
}

func NewJobRepository(db *PostgresDB) *JobRepository {
	return &JobRepository{
		db: db,
	}
}
