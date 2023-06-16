package api

import (
	"E-Commerce_BE/controller"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	userHandler := controller.NewUserHandler(db)
	productHandler := controller.NewProductHandler(db)
	orderHandler := controller.NewOrderHandler(db)
	categoryHandler := controller.NewCategoryHandler(db)

	e.GET("/", controller.ServerOn)

	// Auth route
	authGroup := e.Group("api/v1/auth")
	authGroup.POST("/register", userHandler.AddUser)
	authGroup.POST("/login", userHandler.SignInUser)

	// Management user
	userGroup := e.Group("api/v1/user")
	userGroup.GET("/get-user", userHandler.GetUserById)
	userGroup.GET("/get-all", userHandler.GetAllUser)
	userGroup.PUT("/update/:id", userHandler.UpdateUser)
	userGroup.DELETE("/delete/:id", userHandler.DeleteUser)

	//Category
	categoryGroup := e.Group("api/v1/category")
	categoryGroup.POST("/create", categoryHandler.CreateCategory)

	//Management Product
	productGroup := e.Group("api/v1/product")
	productGroup.POST("/create", productHandler.AddProduct)
	productGroup.GET("/get-product/:id", productHandler.GetProduct)
	productGroup.GET("/all-product", productHandler.GetAllProduct)
	productGroup.DELETE("/delete/:id", productHandler.DeleteProduct)

	//order
	orderGroup := e.Group("api/v1/order")
	orderGroup.POST("/create", orderHandler.OrderProduct)

}
