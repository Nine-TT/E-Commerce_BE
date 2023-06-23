package repository

import (
	"E-Commerce_BE/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	OrderProduct(int, int, int) error
	GetOrder(int) ([]model.Order, error)
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
		UserID:    uint(userID),
		ProductID: uint(productID),
		Quantity:  quantity,
	}).Error
}

func (db *orderRepository) GetOrder(userID int) (listOrders []model.Order, err error) {
	return listOrders, db.db.Where("user_id = ?", userID).Preload(clause.Associations).Find(&listOrders).Error
}
