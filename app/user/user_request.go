package user

import (
	"time"
)

type RequestUser struct {
	Name            string    `json:"name" validate:"required"`
	Address         string    `json:"address" validate:"required"`
	Photo           string    `json:"photo"`
	Email           string    `json:"email" validate:"required,email"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Password        string    `json:"password"`
	Role            UserRole  `json:"role"`
}

type RequestUserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
