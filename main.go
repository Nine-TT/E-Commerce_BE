package main

import (
	"E-Commerce_BE/api"
	"E-Commerce_BE/config/cors"
	DB "E-Commerce_BE/config/db"
	"E-Commerce_BE/model"
	"E-Commerce_BE/util"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Validate input model
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	// ---------------------------------------

	// load file .env
	util.LoadEnv()
	// ---------------------------------------

	// connect db
	db, err := DB.ConnectDB()
	// ---------------------------------------

	// config cors
	cors.SetupCORS(e)
	// ---------------------------------------

	er := db.AutoMigrate(
		model.User{},
		//model.Role{},
		model.Product{},
		model.Category{},
		model.Order{},
	)

	if er != nil {
		return
	}

	if err != nil {
		fmt.Println("Error connect db: ", err)
		return
	} else {
		fmt.Println("connect db success!")
	}

	// init routes
	api.InitRoutes(e, db)
	// ---------------------------------------

	e.Static("/", "product_images")

	e.Logger.Fatal(e.Start(":5000"))
}
