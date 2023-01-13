package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
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

func (r *JobRepository) CreateTransaction(ctx context.Context) (pgx.Tx, error) {
	return r.db.pool.Begin(ctx)
}

func (r *JobRepository) SaveInTransaction(ctx context.Context, tx pgx.Tx, job model.Job) error {
	query := fmt.Sprintf("INSERT INTO %s (id, size, status, created_at, user_id) VALUES ($1, $2, $3, $4, $5)", jobTable)

	if _, err := tx.Exec(ctx, query, job.Id, job.Size, job.Status, job.CreatedAt, job.UserId); err != nil {
		return err
	}

	return nil
}

func (r *JobRepository) FindAll(ctx context.Context, userId int) ([]model.Job, error) {
	query := fmt.Sprintf("SELECT id, size, status, created_at FROM %s WHERE user_id=$1", jobTable)
	rows, err := r.db.pool.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanJobs(rows)
}
