package loan

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/database"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/helper"
	"github.com/labstack/echo/v4"
)

type LoanRoutes struct{}

func (r LoanRoutes) Route() []helper.Route {
	db := database.GetDBinstance()
	db.AutoMigrate(&Loan{})
	loanRepo := NewRepository(db)
	loanService := NewServices(loanRepo)
	loanHandler := NewHandler(loanService)

	return []helper.Route{

		{
			Method:  echo.POST,
			Path:    "/loan",
			Handler: loanHandler.CreateLoan,
		},
		{
			Method:  echo.POST,
			Path:    "/return",
			Handler: loanHandler.Return,
		},
	}
}
