package controller

import (
	"E-Commerce_BE/model"
	"E-Commerce_BE/repository"
	"E-Commerce_BE/util"
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
	Role      string
	jwt.RegisteredClaims
}

func (h *userHandler) AddUser(ctx echo.Context) error {
	var user model.User

	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if errValid := ctx.Validate(user); errValid != nil {
		return errValid
	}

	hasspass, _ := HashPassword(user.Password)
	user.Password = hasspass

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

	if isTrue := CheckPasswordHash(user.Password, dbUser.Password); isTrue == true {

		claims := &jwtCustomClaims{
			dbUser.ID,
			dbUser.FirstName,
			dbUser.LastName,
			dbUser.RoleCode,
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

	id := ctx.Param("id")
	userId := util.GetInteger(id)

	user, err := h.repo.GetUser(userId)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, "User not found!")
	}

	return ctx.JSON(http.StatusOK, user)

	return nil
}

func (h *userHandler) GetAllUser(ctx echo.Context) error {
	user, err := h.repo.GetAllUser()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})

	}
	return ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) UpdateUser(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	user.ID = uint(intID)
	user, err = h.repo.UpdateUser(user, intID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Update user success",
		"user":    user,
	})
}

func (h *userHandler) DeleteUser(ctx echo.Context) error {
	var user model.User
	id := ctx.Param("id")
	intID, _ := strconv.Atoi(id)
	user.ID = uint(intID)
	user, err := h.repo.DeleteUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Delete user success",
	})

}
