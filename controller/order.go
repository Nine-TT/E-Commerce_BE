package controller

import (
	"E-Commerce_BE/model"
	"E-Commerce_BE/repository"
	"E-Commerce_BE/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type orderHandler struct {
	repo repository.OrderRepository
}

func NewOrderHandler(db *gorm.DB) *orderHandler {
	return &orderHandler{
		repo: repository.NewOrderRepository(db),
	}
}

func (h *orderHandler) OrderProduct(ctx echo.Context) error {
	productId := ctx.FormValue("id")
	quantityIDStr := ctx.FormValue("quantity")
	userIDStr := ctx.FormValue("userID")

	proInt := util.GetInteger(productId)
	quanInt := util.GetInteger(quantityIDStr)
	userInt := util.GetInteger(userIDStr)

	if err := h.repo.OrderProduct(userInt, proInt, quanInt); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	} else {
		return ctx.String(http.StatusOK, "Product Successfully ordered")
	}

	//if productIDInt, err := strconv.Atoi(productId); err != nil {
	//	return ctx.JSON(http.StatusBadRequest, echo.Map{
	//		"error": err.Error(),
	//	})
	//} else {
	//	quantityIDStr := ctx.FormValue("quantity")
	//	if quantityID, err := strconv.Atoi(quantityIDStr); err != nil {
	//		return ctx.JSON(http.StatusBadRequest, echo.Map{
	//			"error": err.Error(),
	//		})
	//	} else {
	//		userIDStr := ctx.FormValue("userID")
	//		userIDInt, _ := strconv.Atoi(userIDStr)
	//		if err := h.repo.OrderProduct(userIDInt, productIDInt, quantityID); err != nil {
	//			return ctx.JSON(http.StatusBadRequest, echo.Map{
	//				"error": err.Error(),
	//			})
	//		} else {
	//			return ctx.String(http.StatusOK, "Product Successfully ordered")
	//		}
	//	}
	//}
}

func (h *orderHandler) GetOrderByUserId(ctx echo.Context) error {
	var listOrders []model.Order

	id := ctx.Param("id")
	idInt := util.GetInteger(id)

	listOrders, err := h.repo.GetOrder(idInt)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, listOrders)
}
