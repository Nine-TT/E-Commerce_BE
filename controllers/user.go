package controllers

import (
	"E-Commerce_BE/model"
	"E-Commerce_BE/repositorie"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler interface {
	AddUser(ctx echo.Context)
	GetUser(ctx echo.Context)
	GetAllUser(ctx echo.Context)
	SignInUser(ctx echo.Context)
	UpdateUser(ctx echo.Context)
	DeleteUser(ctx echo.Context)
}

type userHandler struct {
	repo repositorie.UserRepository
}

func NewUserHandler(db *gorm.DB) *userHandler {
	repo := repositorie.NewUserRepository(db)
	return &userHandler{repo: repo}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *userHandler) AddUser(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Misssing!")
	}

	HashPassword(user.PassWord)

	User, err := h.repo.AddUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			return ctx.JSON(http.StatusConflict, echo.Map{
				"message": "Email already exists",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Server Error",
		})
	}

	ctx.JSON(http.StatusOK, echo.Map{
		"message": "Create user success",
		"user":    User,
	})
	return nil
}

func (h *userHandler) GetUserById(ctx echo.Context) error {

	id := ctx.FormValue("id")
	userId, _ := strconv.Atoi(id)

	user, err := h.repo.GetUser(userId)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, "User not found!")
	}

	return ctx.JSON(http.StatusOK, user)

	return nil
}
