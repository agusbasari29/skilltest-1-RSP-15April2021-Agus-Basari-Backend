package payment

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/database"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/helper"
	"github.com/labstack/echo/v4"
)

type PaymentRoutes struct{}

func (r PaymentRoutes) Route() []helper.Route {

	db := database.GetDBinstance()
	db.AutoMigrate(&FinePayment{})
	repo := NewRepository(db)
	services := NewServices(repo)
	handler := NewHandler(services)

	return []helper.Route{

		{
			Method:  echo.GET,
			Path:    "/payment",
			Handler: handler.A,
		},
		{
			Method:  echo.POST,
			Path:    "/payment",
			Handler: handler.A,
		}, {
			Method:  echo.PUT,
			Path:    "/payment",
			Handler: handler.A,
		}, {
			Method:  echo.DELETE,
			Path:    "/payment",
			Handler: handler.A,
		},
	}
}
