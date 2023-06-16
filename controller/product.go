package controller

import (
	"E-Commerce_BE/model"
	"E-Commerce_BE/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(ctx echo.Context) string {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err.Error()
	}

	src, errOpen := file.Open()
	if errOpen != nil {
		return errOpen.Error()
	}
	defer src.Close()

	fileByte, _ := ioutil.ReadAll(src)
	//fileType := http.DetectContentType(fileByte)
	fileName := "product_images/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	filePaths := "E-Commerce/" + fileName

	ioutil.WriteFile(fileName, fileByte, 0777)
	//fileSize := file.Size

	return filePaths
}

type productHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(db *gorm.DB) *productHandler {
	repo := repository.NewProductRepository(db)
	return &productHandler{repo: repo}
}

func (h *productHandler) AddProduct(ctx echo.Context) error {
	var product model.Product
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	//upload image

	file, errFile := ctx.FormFile("file")
	if errFile != nil {
		return errFile
	}

	src, errOpen := file.Open()
	if errOpen != nil {
		return errOpen
	}
	defer src.Close()

	fileByte, _ := io.ReadAll(src)
	fileName := "product_images/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	filePaths := "E-Commerce/" + fileName

	os.WriteFile(fileName, fileByte, 0777)
	//ioutil.WriteFile(fileName, fileByte, 0777)

	product.Image = filePaths

	//--------------

	product, err := h.repo.AddProduct(product)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, echo.Map{
		"message": "create success",
		"product": product,
	})

	return nil
}

func (h *productHandler) GetProduct(ctx echo.Context) error {
	prodStr := ctx.Param("id")
	prodID, err := strconv.Atoi(prodStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	product, err := h.repo.Getproduct(prodID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, product)
}

func (h *productHandler) GetAllProduct(ctx echo.Context) error {
	product, err := h.repo.GetAllproduct()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, product)
}

func (h *productHandler) UpdateProduct(ctx echo.Context) error {

	var product model.Product
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	prodStr := ctx.Param("product")
	prodID, err := strconv.Atoi(prodStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	product.ID = uint(prodID)
	product, err = h.repo.UpdateProduct(product)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, product)

}
func (h *productHandler) DeleteProduct(ctx echo.Context) error {

	var product model.Product
	prodStr := ctx.Param("id")
	prodID, _ := strconv.Atoi(prodStr)
	product.ID = uint(prodID)
	product, err := h.repo.DeleteProduct(product)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Delete product success",
	})
}
