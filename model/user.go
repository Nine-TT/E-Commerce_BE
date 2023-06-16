package model

import (
	"time"
)

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey" `
	FirstName  string    `json:"first_name" validate:"required"`
	LastName   string    `json:"last_name" validate:"required"`
	Address    string    `json:"address" validate:"required"`
	Phone      string    `json:"phone" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	PassWord   string    `json:"pass_word" validate:"required"`
	GenderCode string    `json:"gender_code" validate:"required"`
	RoleCode   string    `json:"role_code" gorm:"default:R2"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	state      bool      `json:"state"`
}

type Role struct {
	Key   string `gorm:"primaryKey" json:"key"`
	Value string `json:"value"`
}

type Gender struct {
	Key   string `gorm:"primaryKey" json:"key"`
	Value string `json:"value"`
}
