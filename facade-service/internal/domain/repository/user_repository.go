package repository

import (
	"context"
	"fmt"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

type UserRepository struct {
	db *PostgresDB
}

func NewUserRepository(db *PostgresDB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByCredentials(ctx context.Context, email, password string) (model.User, error) {
	query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at FROM %s WHERE login_email=$1 AND password_hash=$2", userTable)
	row := r.db.pool.QueryRow(ctx, query, email, password)

	return scanUser(row)
}

func (r *UserRepository) UpdateLastVisit(ctx context.Context, id int) (err error) {
	query := fmt.Sprintf("UPDATE %s SET visited_at=now() WHERE id=$1", userTable)
	_, err = r.db.pool.Exec(ctx, query, id)

	return
}

func (r *UserRepository) FindById(ctx context.Context, id int) (model.User, error) {
	query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at FROM %s WHERE id=$1", userTable)
	row := r.db.pool.QueryRow(ctx, query, id)

	return scanUser(row)
}
