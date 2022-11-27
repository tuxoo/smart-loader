package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

func (r *JobRepository) Save(ctx context.Context, job model.Job) (uuid.UUID, error) {
	var jobId uuid.UUID

	query := fmt.Sprintf("INSERT INTO %s (name, size, status, created_at) VALUES ($1, $2, $3, $4) RETURNING id", jobTable)
	row := r.db.QueryRow(ctx, query, job.Name, job.Size, job.Status, job.CreatedAt)

	if err := row.Scan(&jobId); err != nil {
		return jobId, err
	}

	return jobId, nil
}
