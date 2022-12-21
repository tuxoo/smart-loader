package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

const jobStageDownloadTable = "job_stage_download"

type JobStageDownloadRepository struct {
	db *PostgresDB
}

func NewJobStageDownloadRepository(db *PostgresDB) *JobStageDownloadRepository {
	return &JobStageDownloadRepository{
		db: db,
	}
}

func (r *JobStageDownloadRepository) CreateTransaction(ctx context.Context) (pgx.Tx, error) {
	return r.db.pool.Begin(ctx)
}

func (r *JobStageDownloadRepository) Save(ctx context.Context, jobStageId int, downloadId uuid.UUID) error {
	query := fmt.Sprintf("INSERT INTO %s (job_stage_id, download_id) VALUES ($1, $2)", jobStageDownloadTable)

	if _, err := r.db.pool.Exec(ctx, query, jobStageId, downloadId); err != nil {
		return err
	}

	return nil
}

func (r *JobStageDownloadRepository) SaveInTransaction(ctx context.Context, tx pgx.Tx, jobStageId int, downloadId uuid.UUID) error {
	query := fmt.Sprintf("INSERT INTO %s (job_stage_id, download_id) VALUES ($1, $2)", jobStageDownloadTable)

	if _, err := tx.Exec(ctx, query, jobStageId, downloadId); err != nil {
		return err
	}

	return nil
}
