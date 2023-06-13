package repositorie

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
	GetAllUser() ([]model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository --> returns new user repository
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (db *userRepository) GetUser(id int) (user model.User, err error) {
	return user, db.db.First(&user, id).Error
}

func (db *userRepository) GetByEmail(email string) (user model.User, err error) {
	return user, db.db.First(&user, "email=?", email).Error
}

func (db *userRepository) GetAllUser() (users []model.User, err error) {
	return users, db.db.Find(&users).Error
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

func (db *userRepository) UpdateUser(user model.User) (model.User, error) {
	if err := db.db.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.db.Model(&user).Updates(&user).Error
}

func (db *userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := db.db.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.db.Delete(&user).Error
}
