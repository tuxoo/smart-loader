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

func (r *JobStageRepository) FindAllByJobId(ctx context.Context, jobId uuid.UUID) ([]model.BriefJobStage, error) {
	var stages []model.BriefJobStage

	query := fmt.Sprintf(`
	SELECT id, urls FROM %s WHERE job_id = $1
	`, jobStageTable)
	rows, err := r.db.pool.Query(ctx, query, jobId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var stage model.BriefJobStage

		if err = composeBriefJobStage(&stage, rows); err != nil {
			return nil, err
		}

		stages = append(stages, stage)
	}

	return stages, nil
}

func composeBriefJobStage(stage *model.BriefJobStage, row pgx.Row) error {
	if err := row.Scan(
		&stage.Id,
		&stage.Urls,
	); err != nil {
		return err
	}

	return nil
}
