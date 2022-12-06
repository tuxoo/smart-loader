package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByCredentials(ctx context.Context, email, password string) (*model.User, error) {
	return nil, nil
}
