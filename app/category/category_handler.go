package category

import (
	"net/http"

	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/helper"
	"github.com/labstack/echo/v4"
)

type categoryHandler struct {
	categoryServices Services
}

func NewHandler(categoryServices Services) *categoryHandler {
	return &categoryHandler{categoryServices}
}

func (h *categoryHandler) AddCategory(c echo.Context) error {
	req := new(RequestCategory)
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error1", errorMessage, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	newCat, err := h.categoryServices.AddCategory(*req)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error2", errorMessage, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	catData := CategoryResponseFormatter(newCat)
	response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully add category.", catData)
	return c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) ShowAllCategories(c echo.Context) error {
	categories := h.categoryServices.GetAllCategories()

	response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully add category.", categories)
	return c.JSON(http.StatusOK, response)
}
