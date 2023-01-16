package model

import (
	"time"
)

type SignInDTO struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=5,max=64"`
}

type User struct {
	Id           int       `json:"-" db:"id"`
	Name         string    `json:"name" db:"name" binding:"required"`
	LoginEmail   string    `json:"email" db:"login_email" binding:"required"`
	PasswordHash string    `json:"-" db:"-" binding:"required"`
	RegisteredAt time.Time `json:"registeredAt" db:"registered_at"`
	VisitedAt    time.Time `json:"lastVisitAt" db:"visited_at"`
}
