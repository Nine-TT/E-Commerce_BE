package controller

import (
	"E-Commerce_BE/model"
	"E-Commerce_BE/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

// Get category by id
func (h *categoryHandler) GetCategory(ctx echo.Context) error {
	categoryStr := ctx.Param("id")
	categoruInt, err := strconv.Atoi(categoryStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	category, err := h.repo.GetCategory(categoruInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, category)
}

// Get category by code
func (h *categoryHandler) GetCategoryByCode(ctx echo.Context) error {
	categoryStr := ctx.Param("code")
	category, err := h.repo.GetCategoryByCode(categoryStr)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, category)
}

// Get all category
func (h *categoryHandler) GetAllCategory(ctx echo.Context) error {
	categorys, err := h.repo.GetAllCategory()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, categorys)
}

// Update category by id
func (h *categoryHandler) UpdateCategory(ctx echo.Context) error {
	var category model.Category

	if err := ctx.Bind(&category); err != nil {
		return err
	}

	category, err := h.repo.UpdateCategory(category)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":  "update success",
		"category": category,
	})
}

// Delete category by id
func (h *categoryHandler) DeleteCategory(ctx echo.Context) error {
	categoryStr := ctx.Param("id")
	categoruInt, err := strconv.Atoi(categoryStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	_, errDelete := h.repo.DeleteCategory(categoruInt)
	if errDelete != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Delete success!",
	})
}
