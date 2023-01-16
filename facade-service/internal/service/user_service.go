package service

import (
	"context"
	"github.com/google/uuid"
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
	tokenService ITokenService
}

func NewUserService(
	repository repository.IUserRepository,
	hasher hasher.Hasher,
	tokenManager token_manager.TokenManager,
	tokenService ITokenService,
) *UserService {
	return &UserService{
		repository:   repository,
		hasher:       hasher,
		tokenManager: tokenManager,
		tokenService: tokenService,
	}
}

func (s *UserService) SignIn(ctx context.Context, dto model.SignInDTO) (tokens model.Tokens, err error) {
	user, err := s.repository.FindByCredentials(ctx, dto.Email, s.hasher.HashString(dto.Password))
	if err != nil {
		return
	}

	// TODO: transactions?

	refreshToken, err := s.tokenService.CreateNewToken(ctx, user.Id)
	if err != nil {
		return
	}

	accessToken, err := s.tokenManager.GenerateToken(strconv.Itoa(user.Id))
	if err != nil {
		return
	}

	if err = s.repository.UpdateLastVisit(ctx, user.Id); err != nil {
		return
	}

	return model.Tokens{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func (s *UserService) GetById(ctx context.Context, id int) (model.User, error) {
	return s.repository.FindById(ctx, id)
}

func (s *UserService) RefreshToken(ctx context.Context, userId int, token uuid.UUID) (tokens model.Tokens, err error) {
	refreshToken, err := s.tokenService.RefreshToken(ctx, token)
	if err != nil {
		return
	}

	accessToken, err := s.tokenManager.GenerateToken(strconv.Itoa(userId))
	if err != nil {
		return
	}

	return model.Tokens{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
