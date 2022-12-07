package repository

import (
	"context"
	"fmt"
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
	var user model.User

	query := fmt.Sprintf(`
	SELECT id, name, login_email, registered_at, visited_at FROM %s WHERE login_email=$1 AND password_hash=$2
	`, userTable)
	row := r.db.QueryRow(ctx, query, email, password)

	if err := row.Scan(&user.Id, &user.Name, &user.LoginEmail, &user.RegisteredAt, &user.VisitedAt); err != nil {
		return &user, err
	}

	return &user, nil
}
