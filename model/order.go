package model

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	UserID    uint `json:"user_id" form:"user_id" validate:"required"`
	ProductID uint `json:"product_id" form:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" form:"quantity" validate:"required"`
}
