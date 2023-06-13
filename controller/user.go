package controller

import (
	"E-Commerce_BE/model"
	"E-Commerce_BE/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	repo repository.UserRepository
}

func NewUserHandler(db *gorm.DB) *userHandler {
	repo := repository.NewUserRepository(db)
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

type jwtCustomClaims struct {
	Id        uint
	FirstName string
	LastName  string
	jwt.RegisteredClaims
}

func (h *userHandler) AddUser(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Misssing!")
	}

	hasspass, _ := HashPassword(user.PassWord)
	user.PassWord = hasspass

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

func (h *userHandler) SignInUser(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "Bind data error!")
	}

	dbUser, err := h.repo.GetByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "User not found",
		})
		return nil

	}

	if isTrue := CheckPasswordHash(user.PassWord, dbUser.PassWord); isTrue == true {

		claims := &jwtCustomClaims{
			user.ID,
			user.FirstName,
			user.LastName,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		ctx.JSON(http.StatusOK, echo.Map{
			"message": "login success",
			"token":   t,
		})
		return nil
	}
	ctx.JSON(http.StatusInternalServerError, echo.Map{
		"message": "Password incorrect",
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
