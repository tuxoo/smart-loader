package repository

import (
	"context"
	"fmt"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
)

const downloadTable = "download"

type DownloadRepository struct {
	db *PostgresDB
}

func NewDownloadRepository(db *PostgresDB) *DownloadRepository {
	return &DownloadRepository{
		db: db,
	}
}

func (r *DownloadRepository) FindByHash(ctx context.Context, hash string) (*model.Download, error) {
	var download model.Download

	query := fmt.Sprintf("SELECT id, hash, downloaded_at FROM %s WHERE hash = $1", downloadTable)
	row := r.db.pool.QueryRow(ctx, query, hash)
	if err := row.Scan(&download); err != nil {
		return nil, err
	}

	return &download, nil
}

func (r *DownloadRepository) SaveOne(ctx context.Context, download *model.Download) error {
	query := fmt.Sprintf("INSERT INTO %s (id, hash, downloaded_at) VALUES ($1, $2, $3)", downloadTable)
	if _, err := r.db.pool.Exec(ctx, query, download.Id, download.Hash, download.DownloadedAt); err != nil {
		return err
	}

	return nil
}
