package repository

import (
	"E-Commerce_BE/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	AddUser(user model.User) (model.User, error)
	GetUser(int) (model.User, error)
	GetByEmail(string) (model.User, error)
	GetAllUser(page, pageSize int) ([]model.User, error)
	UpdateUser(model.User, int) (model.User, error)
	DeleteUser(model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository --> returns new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (db *userRepository) GetUser(id int) (user model.User, err error) {
	return user, db.db.Omit("password").Preload("Gender").Preload("Role").Preload("Orders").First(&user, id).Error
}

func (db *userRepository) GetByEmail(email string) (user model.User, err error) {
	return user, db.db.First(&user, "email=?", email).Error
}

func (db *userRepository) CountAllUsers() (count int, err error) {
	var totalCount int64
	err = db.db.Model(&model.User{}).Count(&totalCount).Error
	if err != nil {
		return 0, err
	}
	return int(totalCount), nil
}

func (db *userRepository) GetAllUser(page, pageSize int) (users []model.User, err error) {
	offset := (page - 1) * pageSize
	fmt.Println("============: ", page)
	return users, db.db.Omit("password").Preload("Gender").Preload("Role").Preload("Orders").Offset(offset).Limit(pageSize).Find(&users).Error
}

func (db *userRepository) AddUser(user model.User) (model.User, error) {
	var User model.User
	err := db.db.Where("email = ?", user.Email).First(&User).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ==> Create user
			err = db.db.Create(&user).Error
			return user, err
		}

		return user, err
	}

	// email already exists
	return user, fmt.Errorf("email already exists")
}

func (db *userRepository) UpdateUser(new_user model.User, id int) (user model.User, err error) {
	if err := db.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, db.db.Model(&user).Updates(&new_user).Error
}

func (db *userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := db.db.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.db.Delete(&user).Error
}
