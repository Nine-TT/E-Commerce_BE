package model

import "time"

type Product struct {
	ID          uint      `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Price_sale  float32   `json:"price_sale"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      bool      `json:"status"`
}
