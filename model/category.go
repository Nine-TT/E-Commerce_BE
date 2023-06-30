package model

type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey" form:"id"`
	Code        string    `json:"code" form:"code"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Products    []Product `gorm:"foreignkey:CategoryId"`
}
