package book

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/database"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/helper"
	"github.com/labstack/echo/v4"
)

type BookRoutes struct{}

func (r BookRoutes) Route() []helper.Route {
	db := database.GetDBinstance()
	db.AutoMigrate(&Book{})
	bookRepo := NewRepository(db)
	bookService := NewServices(bookRepo)
	bookHandler := NewHandler(bookService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/books",
			Handler: bookHandler.BookRegister,
		},
		{
			Method:  echo.GET,
			Path:    "/books",
			Handler: bookHandler.GetAllBook,
		},
	}

}
