package repository

import (
	"E-Commerce_BE/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(model.Category) (model.Category, error)
	GetCategory(int) (model.Category, error)
	GetCategoryByCode(string) (model.Category, error)
	GetAllCategory() ([]model.Category, error)
	UpdateCategory(model.Category) (model.Category, error)
	DeleteCategory(int) (model.Category, error)
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

func (db *categoryRepository) GetCategory(id int) (category model.Category, err error) {
	return category, db.db.First(&category, id).Error
}

func (db *categoryRepository) GetCategoryByCode(code string) (category model.Category, err error) {
	return category, db.db.First(&category, "code=?", code).Error
}

func (db *categoryRepository) GetAllCategory() (categorys []model.Category, err error) {
	return categorys, db.db.Find(&categorys).Error
}

func (db *categoryRepository) UpdateCategory(category model.Category) (model.Category, error) {
	if err := db.db.First(&category, category.ID).Error; err != nil {
		return category, err
	}
	return category, db.db.Model(&category).Updates(&category).Error
}

func (db *categoryRepository) DeleteCategory(id int) (category model.Category, err error) {
	if err := db.db.First(&category, "id=?", id).Error; err != nil {
		return category, err
	}
	return category, db.db.Delete(&category).Error
}
