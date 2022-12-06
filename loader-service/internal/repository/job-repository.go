package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/smart-loader/loader-service/internal/model"
)

type JobRepository struct {
	db *pgxpool.Pool
}

func NewJobRepository(db *pgxpool.Pool) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) CreateTransaction(ctx context.Context) (pgx.Tx, error) {
	return r.db.Begin(ctx)
}

func (r *JobRepository) Save(ctx context.Context, tx pgx.Tx, job model.Job) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, size, status, created_at) VALUES ($1, $2, $3, $4, $5)", jobTable)

	if _, err := tx.Exec(ctx, query, job.Id, job.Name, job.Size, job.Status, job.CreatedAt); err != nil {
		return err
	}

	return nil
}
