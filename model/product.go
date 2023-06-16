package model

import (
	"time"
)

type Product struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	CategoryId   uint   `json:"category_id" form:"category_id" validate:"required"`
	Product_code string `json:"product_code" form:"product_code" validate:"required"`
	Title        string `json:"title" form:"title" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
	Image        string
	Price        float32 `json:"price" form:"price" validate:"required"`
	Price_sale   float32 `json:"price_sale" form:"price_sale"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Status       bool `json:"status" form:"code"`
}

type ProductImage struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	ProductId     uint   `json:"product_id"`
	Product_image string `json:"product_images"`
}
