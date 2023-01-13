package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
)

const jobStageTable = "job_stage"

type JobStageRepository struct {
	db *PostgresDB
}

func NewJobStageRepository(db *PostgresDB) *JobStageRepository {
	return &JobStageRepository{
		db: db,
	}
}

func (r *JobStageRepository) CreateTransaction(ctx context.Context) (pgx.Tx, error) {
	return r.db.pool.Begin(ctx)
}

func (r *JobStageRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	query := fmt.Sprintf("UPDATE %s SET status=$1 WHERE id=$2", jobStageTable)
	if _, err := r.db.pool.Exec(ctx, query, status, id); err != nil {
		return err
	}
	return nil
}

func (r *JobStageRepository) UpdateStatusInTransaction(ctx context.Context, tx pgx.Tx, id int, status string) error {
	query := fmt.Sprintf("UPDATE %s SET status=$1 WHERE id=$2", jobStageTable)
	if _, err := tx.Exec(ctx, query, status, id); err != nil {
		return err
	}
	return nil
}

func (r *JobStageRepository) FindAllByJobId(ctx context.Context, jobId uuid.UUID) ([]model.BriefJobStage, error) {
	query := fmt.Sprintf("SELECT id, urls, status FROM %s WHERE job_id = $1", jobStageTable)
	rows, err := r.db.pool.Query(ctx, query, jobId)
	if err != nil {
		return nil, err
	}

	return scanBriefJobStages(rows)
}
