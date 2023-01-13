package repository

import (
	"context"
	"fmt"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
)

const userTable = "\"user\""

type UserRepository struct {
	db *PostgresDB
}

func NewUserRepository(db *PostgresDB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByCredentials(ctx context.Context, email, password string) (*model.User, error) {
	var user model.User

	query := fmt.Sprintf(`
	SELECT id, name, login_email, registered_at, visited_at FROM %s WHERE login_email=$1 AND password_hash=$2
	`, userTable)
	row := r.db.pool.QueryRow(ctx, query, email, password)

	if err := row.Scan(&user.Id, &user.Name, &user.LoginEmail, &user.RegisteredAt, &user.VisitedAt); err != nil {
		return &user, err
	}

	return &user, nil
}
