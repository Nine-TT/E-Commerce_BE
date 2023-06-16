package controller

import (
	"E-Commerce_BE/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
	if productIDInt, err := strconv.Atoi(productId); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	} else {
		quantityIDStr := ctx.FormValue("quantity")
		if quantityID, err := strconv.Atoi(quantityIDStr); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"error": err.Error(),
			})
		} else {
			userIDStr := ctx.FormValue("userID")
			userIDInt, _ := strconv.Atoi(userIDStr)
			if err := h.repo.OrderProduct(userIDInt, productIDInt, quantityID); err != nil {
				return ctx.JSON(http.StatusBadRequest, echo.Map{
					"error": err.Error(),
				})
			} else {
				return ctx.String(http.StatusOK, "Product Successfully ordered")
			}
		}
	}
}
