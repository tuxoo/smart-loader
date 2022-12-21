package model

import (
	"github.com/google/uuid"
	"time"
)

type Download struct {
	Id           uuid.UUID `db:"id"`
	Hash         string    `db:"hash"`
	Size         int       `db:"size"`
	DownloadedAt time.Time `db:"created_at"`
}
