package repository

import (
	"E-Commerce_BE/model"
	"gorm.io/gorm"
	"sort"
)

type ProductRepository interface {
	AddProduct(model.Product) (model.Product, error)
	Getproduct(int) (model.Product, error)
	GetAllproduct() ([]model.Product, error)
	GetProductsByPage(int, int) ([]model.Product, error)
	SortedProductsByPriceMin() ([]model.Product, error)
	SortedProductsByPriceMax() ([]model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(model.Product) (model.Product, error)
	GetSortedProductsByPriceMinAndPage(int, int) ([]model.Product, error)
	GetSortedProductsByPriceMaxAndPage(int, int) ([]model.Product, error)
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

func (db *productRepository) GetProductsByPage(page int, pageSize int) ([]model.Product, error) {
	offset := (page - 1) * pageSize

	products, err := db.GetAllproduct()
	if err != nil {
		return nil, err
	}

	if offset >= len(products) {
		return nil, nil // Kết quả trống nếu trang vượt quá số lượng sản phẩm
	}

	endIndex := offset + pageSize
	if endIndex > len(products) {
		endIndex = len(products)
	}

	return products[offset:endIndex], nil
}

func (db *productRepository) SortedProductsByPriceMin() ([]model.Product, error) {
	products, err := db.GetAllproduct()
	if err != nil {
		return nil, err
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})

	return products, nil
}

func (db *productRepository) SortedProductsByPriceMax() ([]model.Product, error) {
	products, err := db.GetAllproduct()
	if err != nil {
		return nil, err
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})

	return products, nil
}

func (db *productRepository) GetSortedProductsByPriceMinAndPage(page, pageSize int) ([]model.Product, error) {
	products, err := db.SortedProductsByPriceMin()
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize

	if offset >= len(products) {
		return nil, nil
	}

	endIndex := offset + pageSize
	if endIndex > len(products) {
		endIndex = len(products)
	}

	return products[offset:endIndex], nil
}

func (db *productRepository) GetSortedProductsByPriceMaxAndPage(page, pageSize int) ([]model.Product, error) {
	products, err := db.SortedProductsByPriceMax()
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize

	if offset >= len(products) {
		return nil, nil
	}

	endIndex := offset + pageSize
	if endIndex > len(products) {
		endIndex = len(products)
	}

	return products[offset:endIndex], nil
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
