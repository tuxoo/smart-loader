package service

import (
	"github.com/google/uuid"
	"time"
)

type IJobService interface {
	ProcessJob(jobId uuid.UUID) error
}

type IJobStageService interface {
	ProcessStage(jobId uuid.UUID) error
}

type ILockService interface {
	TryToLock(types, value string, expiredAt time.Duration) bool
	TryToUnlock(types, value string) bool
}
