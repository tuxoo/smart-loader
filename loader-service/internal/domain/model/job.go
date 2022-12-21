package model

import (
	"github.com/google/uuid"
	"time"
)

type JobStatus string

type JobStatusDto struct {
	Id        uuid.UUID `json:"id"`
	Status    JobStatus `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type Job struct {
	Id        uuid.UUID `json:"-" db:"id"`
	Name      string    `json:"name" db:"name" binding:"required"`
	Size      int       `json:"size" db:"size" binding:"required"`
	Status    JobStatus `json:"status" db:"status" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" binding:"required"`
}
