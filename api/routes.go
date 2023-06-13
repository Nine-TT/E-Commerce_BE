package api

import (
	"E-Commerce_BE/controller"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	userHandler := controller.NewUserHandler(db)
	e.GET("/", controller.ServerOn)

	// Auth route
	authGroup := e.Group("api/v1/auth")
	authGroup.POST("/register", userHandler.AddUser)
	authGroup.POST("/login", userHandler.SignInUser)

	// Management user
	userGroup := e.Group("api/v1/user")
	userGroup.GET("/get-user", userHandler.GetUserById)

}
