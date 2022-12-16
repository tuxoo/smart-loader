package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

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

func (r *JobRepository) Save(ctx context.Context, tx pgx.Tx, job model.Job) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, size, status, created_at, user_id) VALUES ($1, $2, $3, $4, $5, $6)", jobTable)

	if _, err := tx.Exec(ctx, query, job.Id, job.Name, job.Size, job.Status, job.CreatedAt, job.UserId); err != nil {
		return err
	}

	return nil
}
