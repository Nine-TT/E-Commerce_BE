package main

import (
	"E-Commerce_BE/api"
	DB "E-Commerce_BE/config/db"
	"E-Commerce_BE/model"
	"E-Commerce_BE/util"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Path string
}

func upload(ctx echo.Context) error {
	var response Response

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileByte, _ := ioutil.ReadAll(src)
	//fileType := http.DetectContentType(fileByte)
	fileName := "product_images/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	filePaths := "E-Commerce/" + fileName

	ioutil.WriteFile(fileName, fileByte, 0777)
	//fileSize := file.Size

	fmt.Println("+=======>>> path", filePaths)
	response.Path = filePaths

	return ctx.JSON(http.StatusOK, echo.Map{
		"path": response.Path,
	})
}

//func uploadMultiple(ctx echo.Context) error {
//	var response []Response
//
//	form, _ := ctx.MultipartForm()
//	files, _ := form.File["files"]
//
//	for i := 0; i < len(files); i++ {
//		src, err := files[i].Open()
//		if err != nil {
//			return err
//		}
//		defer src.Close()
//
//		fileByte, _ := ioutil.ReadAll(src)
//		fileType := http.DetectContentType(fileByte)
//		fileName := "product_images/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
//
//
//
//		fmt.Printf("src type %T", src)
//	}
//	//for _, file := range files {
//	//	src, _ := file.Open()
//	//	defer src.Close()
//	//}
//
//	return ctx.JSON(http.StatusOK, echo.Map{
//		"response": response,
//	})
//}

func main() {
	e := echo.New()

	e.Validator = &util.CustomValidator{Validator: validator.New()}

	util.LoadEnv()

	db, err := DB.ConnectDB()

	er := db.AutoMigrate(
		model.User{},
		model.Product{},
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

	e.Static("/", "product_images")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":5000"))
}
