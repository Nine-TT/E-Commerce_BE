package repository

import (
	"E-Commerce_BE/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(model.Category) (model.Category, error)
	//GetCategory(int) (model.Category , error)
	//GetAllCategory() ([]model.Category, error)
	//UpdateCategory( model.Category) (model.Category, error)
	//DeleteCategory( model.Category) (model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (db *categoryRepository) CreateCategory(category model.Category) (model.Category, error) {
	return category, db.db.Create(&category).Error
}
