package model

import (
	"github.com/google/uuid"
	"time"
)

type Download struct {
	Id           uuid.UUID `db:"id"`
	Hash         string    `db:"hash"`
	DownloadedAt time.Time `db:"downloaded_at"`
	Size         int       `db:"size"`
}
