package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
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

func (r *DownloadRepository) CreateTransaction(ctx context.Context) (pgx.Tx, error) {
	return r.db.pool.Begin(ctx)
}

func (r *DownloadRepository) FindByHash(ctx context.Context, hash string) (*model.Download, error) {
	query := fmt.Sprintf("SELECT id, hash, downloaded_at, size FROM %s WHERE hash = $1", downloadTable)
	rows, err := r.db.pool.Query(ctx, query, hash)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var download model.Download

		if err = rows.Scan(&download.Id, &download.Hash, &download.DownloadedAt, &download.Size); err != nil {
			return nil, err
		}

		return &download, nil
	}

	return nil, nil
}

func (r *DownloadRepository) SaveOne(ctx context.Context, tx pgx.Tx, download *model.Download) error {
	query := fmt.Sprintf("INSERT INTO %s (id, hash, size, downloaded_at) VALUES ($1, $2, $3, $4)", downloadTable)
	if _, err := tx.Exec(ctx, query, download.Id, download.Hash, download.Size, download.DownloadedAt); err != nil {
		return err
	}

	return nil
}
