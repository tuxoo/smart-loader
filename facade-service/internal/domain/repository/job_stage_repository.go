package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
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

func (r *JobStageRepository) SaveInTransaction(ctx context.Context, tx pgx.Tx, jobStage model.JobStage) error {
	query := fmt.Sprintf("INSERT INTO %s (size, urls, status, job_id) VALUES ($1, $2, $3, $4)", jobStageTable)

	if _, err := tx.Exec(ctx, query, jobStage.Size, jobStage.Urls, jobStage.Status, jobStage.JobId); err != nil {
		return err
	}

	return nil
}

func (r *JobStageRepository) FindAllByJobId(ctx context.Context, jobId uuid.UUID) ([]int, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE job_id = $1", jobStageTable)

	rows, err := r.db.pool.Query(ctx, query, jobId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int

		if err = rows.Scan(&id); err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}
