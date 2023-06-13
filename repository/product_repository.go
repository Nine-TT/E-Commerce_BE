package repository

import "E-Commerce_BE/model"

type ProductRepository interface {
	Getproduct(int) (model.Product, error)
	GetAllproduct() ([]model.Product, error)
	AddProduct(model.Product) (model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(model.Product) (model.Product, error)
}
