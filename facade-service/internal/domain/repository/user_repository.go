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
	query := fmt.Sprintf(`
	SELECT id, name, login_email, registered_at, visited_at FROM %s WHERE login_email=$1 AND password_hash=$2
	`, userTable)
	row := r.db.pool.QueryRow(ctx, query, email, password)

	return scanUser(row)
}
