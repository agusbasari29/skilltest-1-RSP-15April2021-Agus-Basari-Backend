package book

import (
	"net/http"

	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/helper"
	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	bookService Services
}

func NewHandler(bookService Services) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) BookRegister(c echo.Context) error {
	req := new(RequestBook)

	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	existBook := h.bookService.IsBookExist(*req)
	if existBook != nil {
		if err := h.bookService.UpdateStock(1, *req); err != nil {
			errorFormatter := helper.ErrorFormatter(err)
			errorMessage := helper.M{"errors": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			return c.JSON(http.StatusBadRequest, response)
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Book is exist, adding stock only.", nil)
		return c.JSON(http.StatusOK, response)
	} else {
		newBook, err := h.bookService.CreateBook(*req)
		if err != nil {
			errorFormatter := helper.ErrorFormatter(err)
			errorMessage := helper.M{"errors": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			return c.JSON(http.StatusBadRequest, response)
		}
		bookData := BookResponseFormatter(newBook)
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully add category.", bookData)
		return c.JSON(http.StatusOK, response)
	}
}

func (h *bookHandler) GetAllBook(c echo.Context) error {
	allBook := h.bookService.AllBooks()
	response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully add category.", allBook)
	return c.JSON(http.StatusOK, response)
}
