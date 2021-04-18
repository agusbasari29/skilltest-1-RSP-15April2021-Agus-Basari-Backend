package payment

import "github.com/labstack/echo/v4"

type paymentHandler struct {
	paymentServices Services
}

func NewHandler(paymentServices Services) *paymentHandler {
	return &paymentHandler{paymentServices}
}

func (h paymentHandler) A(c echo.Context) error {
	return nil
}
