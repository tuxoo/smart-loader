package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

const jobTable = "job"

type JobRepository struct {
	db *PostgresDB
}

func NewJobRepository(db *PostgresDB) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	query := fmt.Sprintf("UPDATE %s SET status=$1 WHERE id=$2", jobTable)
	if _, err := r.db.pool.Exec(ctx, query, status, id); err != nil {
		return err
	}
	return nil
}
