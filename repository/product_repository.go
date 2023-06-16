package repository

import (
	"E-Commerce_BE/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(model.Product) (model.Product, error)
	Getproduct(int) (model.Product, error)
	GetAllproduct() ([]model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(model.Product) (model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (db *productRepository) AddProduct(product model.Product) (model.Product, error) {
	return product, db.db.Create(&product).Error
}

func (db *productRepository) Getproduct(id int) (product model.Product, err error) {
	return product, db.db.First(&product, id).Error
}

func (db *productRepository) GetAllproduct() (products []model.Product, err error) {
	return products, db.db.Find(&products).Error
}

func (db *productRepository) UpdateProduct(product model.Product) (model.Product, error) {
	if err := db.db.First(&product, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.db.Model(&product).Updates(&product).Error
}

func (db *productRepository) DeleteProduct(product model.Product) (model.Product, error) {
	if err := db.db.First(&product, product.ID).Error; err != nil {
		return product, err
	}
	return product, db.db.Delete(&product).Error
}
