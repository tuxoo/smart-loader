package model

import (
	"github.com/google/uuid"
	"time"
)

type JobStatus string

const (
	NEW     = "NEW"
	PENDING = "PENDING"
)

type Job struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Size      int       `json:"size" db:"size" binding:"required"`
	Status    JobStatus `json:"status" db:"status" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" binding:"required"`
	UserId    int       `json:"-"`
}
