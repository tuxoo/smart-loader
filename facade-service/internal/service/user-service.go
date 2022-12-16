package service

import (
	"context"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/facade-service/internal/util/hasher"
	"github.com/tuxoo/smart-loader/facade-service/internal/util/token-manager"
	"strconv"
)

type UserService struct {
	repository   repository.IUserRepository
	hasher       hasher.Hasher
	tokenManager token_manager.TokenManager
}

func NewUserService(repository repository.IUserRepository, hasher hasher.Hasher, tokenManager token_manager.TokenManager) *UserService {
	return &UserService{
		repository:   repository,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}

func (s *UserService) SignIn(ctx context.Context, dto model.SignInDTO) (token string, err error) {
	user, err := s.repository.FindByCredentials(ctx, dto.Email, s.hasher.Hash(dto.Password))
	if err != nil {
		return
	}

	return s.tokenManager.GenerateToken(strconv.Itoa(user.Id))
}
