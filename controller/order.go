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
	var order model.Order

	if err := ctx.Bind(&order); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"Error": err.Error(),
		})
	}

	if _, err := h.repo.OrderProduct(order); err != nil {
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

	id := util.GetInteger(ctx.Param("id"))

	listOrders, err := h.repo.GetOrder(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, listOrders)
}

func (h *orderHandler) UpdateOrderQuanty(ctx echo.Context) error {
	id := util.GetInteger(ctx.Param("id"))
	var order model.Order

	ctx.Bind(&order)

	order, err := h.repo.UpdateOrderQuantity(order, id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"Error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"mesage": "update quantity success!",
		"order":  order,
	})
}
