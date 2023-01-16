package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

type TokenRepository struct {
	db *PostgresDB
}

func NewTokenRepository(db *PostgresDB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) SaveOne(ctx context.Context, token model.Token) (id uuid.UUID, err error) {
	query := fmt.Sprintf("INSERT INTO %s (expired_at, user_id) VALUES ($1, $2) RETURNING id", tokenTable)
	row := r.db.pool.QueryRow(ctx, query, token.ExpiredAt, token.UserId)

	if err = row.Scan(&id); err != nil {
		return
	}

	return
}

func (r *TokenRepository) FindAllByUser(ctx context.Context, userId int) ([]model.Token, error) {
	query := fmt.Sprintf("SELECT id, expired_at, user_id FROM %s WHERE user_id=$1", tokenTable)
	rows, err := r.db.pool.Query(ctx, query, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTokens(rows)
}

func (r *TokenRepository) DeleteByUser(ctx context.Context, userId int) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", tokenTable)
	_, err = r.db.pool.Exec(ctx, query, userId)

	return
}

func (r *TokenRepository) UpdateToken(ctx context.Context, token model.Token) (id uuid.UUID, err error) {
	query := fmt.Sprintf("UPDATE %s SET id=gen_random_uuid(), expired_at=$1 WHERE id = $2 RETURNING id", tokenTable)
	row := r.db.pool.QueryRow(ctx, query, token.ExpiredAt, token.Id)

	if err = row.Scan(&id); err != nil {
		return
	}

	return
}
