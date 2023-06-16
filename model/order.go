package model

type Order struct {
	ID        uint    `json:"id"`
	User      User    `gorm:"foreignkey:UserID"`
	Product   Product `gorm:"foreignkey:ProductID"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
}
