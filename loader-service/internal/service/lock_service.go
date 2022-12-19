package service

import (
	"context"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
)

type LockService struct {
	repository repository.ILockRepository
}

func NewLockService(repository repository.ILockRepository) *LockService {
	return &LockService{
		repository: repository,
	}
}

func (s *LockService) TryToLock(ctx context.Context, types, value string) bool {
	err := s.repository.Lock(ctx, types, value)
	if err != nil {
		return false
	}

	return true
}

func (s *LockService) TryToUnlock(ctx context.Context, types, value string) bool {
	err := s.repository.Unlock(ctx, types, value)
	if err != nil {
		return false
	}

	return true
}
