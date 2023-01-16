package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"
	"time"
)

type TokenService struct {
	cfg        *config.AppConfig
	repository repository.ITokenRepository
}

func NewTokenService(cfg *config.AppConfig, repository repository.ITokenRepository) *TokenService {
	return &TokenService{
		cfg:        cfg,
		repository: repository,
	}
}

func (s *TokenService) CreateNewToken(ctx context.Context, userId int) (token uuid.UUID, err error) {
	tokens, err := s.repository.FindAllByUser(ctx, userId)
	if err != nil {
		return
	}

	if len(tokens) >= s.cfg.TokensLimit {
		if err = s.repository.DeleteByUser(ctx, userId); err != nil {
			return
		}
	}

	return s.repository.SaveOne(ctx, model.Token{
		ExpiredAt: time.Now().Add(s.cfg.RefreshTokenTTL),
		UserId:    userId,
	})
}

func (s *TokenService) RefreshToken(ctx context.Context, token uuid.UUID) (uuid.UUID, error) {
	return s.repository.UpdateToken(ctx, model.Token{
		Id:        token,
		ExpiredAt: time.Now().Add(s.cfg.RefreshTokenTTL),
	})
}
