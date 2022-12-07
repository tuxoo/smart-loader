package service

import (
	"context"
	"fmt"
	"github.com/tuxoo/smart-loader/facade-service/internal/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
	"github.com/tuxoo/smart-loader/facade-service/internal/util"
)

type UserService struct {
	repository repository.IUserRepository
	hasher     *util.Hasher
}

func NewUserService(repository repository.IUserRepository, hasher *util.Hasher) *UserService {
	return &UserService{
		repository: repository,
		hasher:     hasher,
	}
}

func (s *UserService) SignIn(ctx context.Context, dto model.SignInDTO) (token string, err error) {
	user, err := s.repository.FindByCredentials(ctx, dto.Email, s.hasher.SHA256Hash(dto.Password))
	if err != nil {
		return
	}

	fmt.Println(user)
	return
}
