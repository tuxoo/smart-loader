package repository

import (
	"context"
	"fmt"
	"time"
)

const lockTable = "lock"

type LockRepository struct {
	db *PostgresDB
}

func NewLockRepository(db *PostgresDB) *LockRepository {
	return &LockRepository{
		db: db,
	}
}

func (r *LockRepository) Lock(ctx context.Context, types, value string) error {
	query := fmt.Sprintf(`
	INSERT INTO %s (type, value, expired_at)
	VALUES ($1, $2, $3)
	`, lockTable)

	if _, err := r.db.pool.Exec(ctx, query, types, value, time.Now()); err != nil {
		return err
	}

	return nil
}

func (r *LockRepository) Unlock(ctx context.Context, types, value string) error {
	query := fmt.Sprintf("DELETE FROM %s where type =$1 and value = $2", lockTable)

	if _, err := r.db.pool.Exec(ctx, query, types, value); err != nil {
		return err
	}

	return nil
}
