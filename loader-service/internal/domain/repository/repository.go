package repository

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
	ChangeState(types, value string, state bool) error
}
