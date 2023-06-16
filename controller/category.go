package controller

import (
	"E-Commerce_BE/model"
	"E-Commerce_BE/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type categoryHandler struct {
	repo repository.CategoryRepository
}

func NewCategoryHandler(db *gorm.DB) *categoryHandler {
	repo := repository.NewCategoryRepository(db)
	return &categoryHandler{repo: repo}
}

func (h *categoryHandler) CreateCategory(ctx echo.Context) error {
	var category model.Category
	if err := ctx.Bind(&category); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	category, err := h.repo.CreateCategory(category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, echo.Map{
		"message":  "create success",
		"category": category,
	})

	return nil
}
