package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

type DownloadRepository struct {
	db *PostgresDB
}

func NewDownloadRepository(db *PostgresDB) *DownloadRepository {
	return &DownloadRepository{
		db: db,
	}
}

func (r *DownloadRepository) FindAllByJobId(ctx context.Context, jobId uuid.UUID, userId int) ([]model.Download, error) {
	query := fmt.Sprintf(`
	SELECT d.id, d.hash, d.downloaded_at, d.size
	FROM %s
			 INNER JOIN %s js ON job.id = js.job_id
			 INNER JOIN %s jsd ON js.id = jsd.job_stage_id
			 INNER JOIN %s d ON d.id = jsd.download_id
	WHERE job.id = $1 AND job.user_id = $2
	`, jobTable, jobStageTable, jobStageDownloadTable, downloadTable)

	rows, err := r.db.pool.Query(ctx, query, jobId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanDownloads(rows)
}
