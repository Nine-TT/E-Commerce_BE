package main

import (
	"E-Commerce_BE/api"
	DB "E-Commerce_BE/config/db"
	"E-Commerce_BE/model"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := DB.ConnectDB()

	er := db.AutoMigrate(
		model.User{},
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

	api.InitRoutes(e, db)

	e.Logger.Fatal(e.Start(":5000"))
}
