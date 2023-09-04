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
	userGroup.GET("/:id", userHandler.GetUserById)
	//userGroup.Use(middleware.IsAdmin())
	userGroup.GET("/all", userHandler.GetAllUser)
	userGroup.PUT("/update/:id", userHandler.UpdateUser)
	userGroup.DELETE("/delete/:id", userHandler.DeleteUser)

	//Category
	//admin + manager (Role)
	categoryGroup := e.Group("api/v1/category")
	//categoryGroup.Use(middleware.IsAdmin(), middleware.IsManagement())
	categoryGroup.POST("/create", categoryHandler.CreateCategory)
	categoryGroup.GET("/:id", categoryHandler.GetCategory)
	categoryGroup.GET("/get-by-code/:code", categoryHandler.GetCategoryByCode)
	categoryGroup.GET("/all-category", categoryHandler.GetAllCategory)
	categoryGroup.PUT("/update", categoryHandler.UpdateCategory)
	categoryGroup.DELETE("/delete/:id", categoryHandler.DeleteCategory)

	//Management Product
	// Admin + manager
	productGroup := e.Group("api/v1/product")
	productGroup.POST("/create", productHandler.AddProduct)
	productGroup.GET("/get-product/:id", productHandler.GetProduct)
	// productGroup.GET("/list", productHandler.GetListProducts)
	productGroup.GET("/all-product", productHandler.GetAllProduct)
	productGroup.GET("/product-page", productHandler.GetProductsByPage)
	productGroup.DELETE("/delete/:id", productHandler.DeleteProduct)
	// productGroup.DELETE("/delete/list", productHandler.DeleteProducts)

	// => Filter product
	productGroup.GET("/price/min", productHandler.SortProductPriceMin)
	productGroup.GET("/price/max", productHandler.SortProductPriceMax)
	productGroup.GET("/price/between", productHandler.ProductTwoPrice)
	productGroup.GET("/search/:name", productHandler.SearchProductName)

	//order
	orderGroup := e.Group("api/v1/order")
	orderGroup.POST("/create", orderHandler.OrderProduct)
	orderGroup.GET("/user/:id", orderHandler.GetOrderByUserId)
	orderGroup.PUT("/update/:id", orderHandler.UpdateOrderQuanty)

	// update order

}
