package model

import (
	"github.com/google/uuid"
)

type JobStageStatus string

type JobStage struct {
	Id     int       `json:"-" db:"id"`
	Size   int       `json:"size" db:"size" binding:"required"`
	Urls   []string  `json:"urls" db:"urls" binding:"required"`
	Status JobStatus `json:"status" db:"status" binding:"required"`
	JobId  uuid.UUID `json:"jobId" db:"job_id" binding:"required"`
}

type BriefJobStage struct {
	Id     int      `json:"-" db:"id"`
	Status string   `json:"status" db:"status"`
	Urls   []string `json:"urls" db:"urls" binding:"required"`
}
