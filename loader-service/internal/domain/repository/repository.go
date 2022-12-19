package repository

import "context"

const (
	lockTable     = "lock"
	jobTable      = "job"
	jobStageTable = "job_stage"
	downloadTable = "download"
)

type IJobRepository interface {
}

type IJobStageRepository interface {
}

type ILockRepository interface {
	Lock(ctx context.Context, types, value string) error
	Unlock(ctx context.Context, types, value string) error
}
