package service

import (
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
	"time"
)

type LockService struct {
	repository repository.ILockRepository
}

func NewLockService(repository repository.ILockRepository) *LockService {
	return &LockService{
		repository: repository,
	}
}

func (s *LockService) TryToLock(types, value string, expiredAt time.Duration) bool {
	return false
}

func (s *LockService) TryToUnlock(types, value string) bool {
	return false
}
