package repository

import (
	"E-Commerce_BE/model"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	OrderProduct(model.Order) (model.Order, error)
	GetOrder(int) ([]model.Order, error)
	UpdateOrderQuantity(model.Order, int) (model.Order, error)
	DeleteOrder(model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (db *orderRepository) OrderProduct(new_order model.Order) (model.Order, error) {
	var order model.Order
	err := db.db.Where("user_id = ? AND product_id = ?", new_order.UserID, new_order.ProductID).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create new order
			err = db.db.Create(&new_order).Error
		}

		return order, err
	}

	// order already exists ==> update quantity
	return order, db.db.Model(&order).Where("user_id = ? AND product_id = ?", order.UserID, order.ProductID).Updates(&model.Order{
		UserID:    order.UserID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity + new_order.Quantity,
	}).Error
}

func (db *orderRepository) GetOrder(userID int) (listOrders []model.Order, err error) {
	return listOrders, db.db.Where("user_id = ?", userID).Preload(clause.Associations).Find(&listOrders).Error
}

func (db *orderRepository) UpdateOrderQuantity(new_order model.Order, id int) (model.Order, error) {
	var order model.Order
	if err := db.db.First(&order, id).Error; err != nil {
		return order, err
	}
	return order, db.db.Model(&order).Where("ID = ?", id).Updates(&new_order).Error
}

func (db *orderRepository) DeleteOrder(order model.Order) error {
	return db.db.Delete(&order, order.ID).Error
}
