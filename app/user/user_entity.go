package user

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	Admin  UserRole = `admin`
	Member UserRole = `member`
)

type User struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	Name            string
	Address         string
	Photo           string
	Email           string
	EmailVerifiedAt time.Time
	Password        string
	Role            UserRole
}
