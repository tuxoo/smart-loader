package model

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	Id        uuid.UUID `db:"id"`
	ExpiredAt time.Time `db:"expired_at"`
	UserId    int       `db:"user_id"`
}

type Tokens struct {
	RefreshToken uuid.UUID `json:"refreshToken"`
	AccessToken  string    `json:"accessToken"`
}
