package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
)

type JobStageRepository struct {
	db *pgxpool.Pool
}

func NewJobStageRepository(db *pgxpool.Pool) *JobStageRepository {
	return &JobStageRepository{
		db: db,
	}
}

func (r *JobStageRepository) Save(ctx context.Context, jobStage model.JobStage) error {
	query := fmt.Sprintf("INSERT INTO %s (size, uris, status, job_id) VALUES ($1, $2, $3, $4)", jobStageTable)
	_, err := r.db.Exec(ctx, query, jobStage.Size, jobStage.Uris, jobStage.Status, jobStage.JobId)

	return err
}
