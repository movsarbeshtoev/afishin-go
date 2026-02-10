package models

import (
	"time"
)

const (
	RoleUser = "user"
	RoleModerator = "moderator"
	RoleAdmin = "admin"
)


type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Email string `gorm:"uniqueIndex" json:"email"`
	PasswordHash string `gorm:"column:password_hash;not null" json:"-"`
	Phone string `gorm:"uniqueIndex" json:"phone"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Role         string    `gorm:"default:user;size:20" json:"role"`
}

