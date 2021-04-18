package loan

import (
	"net/http"

	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/helper"
	"github.com/labstack/echo/v4"
)

type loanHandler struct {
	loanServices Services
	// bookService  book.Services
	// userService  user.Services
}

func NewHandler(loanServices Services) *loanHandler {
	return &loanHandler{loanServices}
}

func (h *loanHandler) CreateLoan(c echo.Context) error {
	req := new(RequestLoan)
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	newLoan, err := h.loanServices.CreateLoan(*req)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	loanData := LoanResponseFormatter(newLoan)
	response := helper.ResponseFormatter(http.StatusOK, "success", "Loan successfully created.", loanData)
	return c.JSON(http.StatusOK, response)
}

func (h *loanHandler) Return(c echo.Context) error {
	req := new(RequestLoan)
	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	loan := h.loanServices.GetLoan(*req)

	loanData := LoanResponseFormatter(*loan)
	response := helper.ResponseFormatter(http.StatusOK, "success", "Loan successfully created.", loanData)
	return c.JSON(http.StatusOK, response)
}
