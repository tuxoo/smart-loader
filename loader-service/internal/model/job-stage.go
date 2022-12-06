package model

import (
	"github.com/google/uuid"
)

type JobStageStatus string

type JobStage struct {
	Id     uuid.UUID `json:"-" db:"id"`
	Size   int       `json:"size" db:"size" binding:"required"`
	Uris   []string  `json:"uris" db:"uris" binding:"required"`
	Status JobStatus `json:"status" db:"status" binding:"required"`
	JobId  uuid.UUID `json:"jobId" db:"job_id" binding:"required"`
}
