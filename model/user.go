package model

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Address    string    `json:"address"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	PassWord   string    `json:"pass_word"`
	GenderCode string    `json:"gender_code"`
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
