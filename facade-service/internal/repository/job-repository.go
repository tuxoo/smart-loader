package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
)

type JobRepository struct {
	db *pgxpool.Pool
}

func NewJobRepository(db *pgxpool.Pool) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) Save(ctx context.Context, job model.Job) error {
	query := fmt.Sprintf("INSERT INTO %s (name, size, status, created_at) VALUES ($1, $2, $3, $4)", jobTable)
	_, err := r.db.Exec(ctx, query, job.Name, job.Size, job.Status, job.CreatedAt)

	return err
}
