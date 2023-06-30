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
	Password   string    `json:"password" validate:"required"`
	GenderCode string    `json:"gender_code" gorm:"varchar(50)" validate:"required"`
	RoleCode   string    `json:"role_code" gorm:"Type:varchar(50); default:R3"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	state      bool      `json:"state"`

	//belongs to
	Role   Role   `gorm:"foreignkey:RoleCode"`
	Gender Gender `gorm:"foreignkey:GenderCode"`

	//has many (one-to-many)
	Orders []Order `gorm:"foreignkey:UserID"`
}

type Role struct {
	Key   string `gorm:"primaryKey; Type:varchar(50)" json:"key"`
	Value string `json:"value"`
}

type Gender struct {
	Key   string `gorm:"primaryKey; Type:varchar(50)" json:"key"`
	Value string `json:"value"`
}
