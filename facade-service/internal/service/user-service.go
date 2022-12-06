package service

import (
	"context"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) SignIn(ctx context.Context, dto model.SignInDTO) (string, error) {
	return "", nil
}
