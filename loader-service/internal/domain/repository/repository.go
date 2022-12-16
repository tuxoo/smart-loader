package repository

const (
	userTable     = "\"user\""
	jobTable      = "job"
	jobStageTable = "job_stage"
	downloadTable = "download"
)

type IJobRepository interface {
}

type IJobStageRepository interface {
}
