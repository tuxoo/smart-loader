package service

import (
	"context"
	"github.com/google/uuid"
)

type IJobService interface {
}

type IJobStageService interface {
	ProcessStage(ctx context.Context, jobId uuid.UUID) error
}

type ILockService interface {
	TryToLock(ctx context.Context, types, value string) bool
	TryToUnlock(ctx context.Context, types, value string) bool
}
