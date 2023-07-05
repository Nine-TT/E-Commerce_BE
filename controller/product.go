package controller

import (
	"E-Commerce_BE/model"
	"E-Commerce_BE/repository"
	"E-Commerce_BE/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
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

	path := "product_images/" + product.Product_code

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0660)
	}

	fileByte, _ := io.ReadAll(src)
	fileName := path + "/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
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
	prodID := util.GetInteger(ctx.Param("id"))
	product, err := h.repo.Getproduct(prodID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, product)
}

func (h *productHandler) GetListProducts(ctx echo.Context) error {
	var listIDs struct {
		IDs []int `json:"id"`
	}

	if err := ctx.Bind(&listIDs); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error Bind Data!")
	}

	listPodcuts, err := h.repo.GetProdcuts(listIDs.IDs)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, echo.Map{
			"Error": err.Error(),
		})
		return err
	}

	return ctx.JSON(http.StatusOK, listPodcuts)
}

func (h *productHandler) GetAllProduct(ctx echo.Context) error {
	product, err := h.repo.GetAllproduct()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, product)
}

func (h *productHandler) GetProductsByPage(ctx echo.Context) error {
	page := util.GetInteger(ctx.FormValue("page"))
	pageSize := util.GetInteger(ctx.FormValue("size"))

	var listProducts []model.Product

	listProducts, err := h.repo.GetProductsByPage(page, pageSize)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"Error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, listProducts)
}

// filter with product price

func (h *productHandler) SortProductPriceMin(ctx echo.Context) error {
	page := util.GetInteger(ctx.FormValue("page"))
	size := util.GetInteger(ctx.FormValue("size"))

	var listProducts []model.Product

	listProducts, err := h.repo.GetSortedProductsByPriceMinAndPage(page, size)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, listProducts)
}

func (h *productHandler) SortProductPriceMax(ctx echo.Context) error {
	page := util.GetInteger(ctx.FormValue("page"))
	size := util.GetInteger(ctx.FormValue("size"))

	var listProducts []model.Product

	listProducts, err := h.repo.GetSortedProductsByPriceMaxAndPage(page, size)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, listProducts)

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

func (h *productHandler) DeleteProducts(ctx echo.Context) error {
	var listIDs struct {
		IDs []int `json:"id"`
	}

	if err := ctx.Bind(&listIDs); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error Bind Data!")
	}

	err := h.repo.DeleteProducts(listIDs.IDs)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"Error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Delete list product success",
	})
}

func (h *productHandler) SearchProductName(ctx echo.Context) error {
	name := ctx.Param("name")

	products, err := h.repo.SearchProducts(name)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"Error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, products)
}

// loc san pham theo gia khoang x - y

func (h *productHandler) ProductTwoPrice(ctx echo.Context) error {
	var prices struct {
		FirstPrice  float32 `json:"first_price"`
		SecondPrice float32 `json:"second_price"`
	}

	if err := ctx.Bind(&prices); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"Error": err.Error()})
	}

	listProducts, _ := h.repo.ProductTwoPrice(prices.FirstPrice, prices.SecondPrice)

	sort.Slice(listProducts, func(i, j int) bool {
		return listProducts[i].Price < listProducts[j].Price
	})

	return ctx.JSON(http.StatusOK, listProducts)
}
