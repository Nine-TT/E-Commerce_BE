package repository

import (
	"E-Commerce_BE/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	OrderProduct(int, int, int) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (db *orderRepository) OrderProduct(userID int, productID int, quantity int) error {
	return db.db.Create(&model.Order{
		ProductID: uint(productID),
		UserID:    uint(userID),
		Quantity:  quantity,
	}).Error
}
