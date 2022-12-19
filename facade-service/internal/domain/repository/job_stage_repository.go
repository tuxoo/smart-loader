package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

type JobStageRepository struct {
	db *PostgresDB
}

func NewJobStageRepository(db *PostgresDB) *JobStageRepository {
	return &JobStageRepository{
		db: db,
	}
}

func (r *JobStageRepository) Save(ctx context.Context, tx pgx.Tx, jobStage model.JobStage) error {
	query := fmt.Sprintf("INSERT INTO %s (size, uris, status, job_id) VALUES ($1, $2, $3, $4)", jobStageTable)

	if _, err := tx.Exec(ctx, query, jobStage.Size, jobStage.Uris, jobStage.Status, jobStage.JobId); err != nil {
		return err
	}

	return nil
}
