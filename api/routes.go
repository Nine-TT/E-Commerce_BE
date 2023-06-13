package api

import (
	"E-Commerce_BE/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	userHandler := controllers.NewUserHandler(db)
	e.GET("/", controllers.ServerOn)

	// Auth route
	authGroup := e.Group("api/v1/auth")
	authGroup.POST("/register", userHandler.AddUser)

	// Management user
	userGroup := e.Group("api/v1/user")
	userGroup.GET("/get-user", userHandler.GetUserById)
}
