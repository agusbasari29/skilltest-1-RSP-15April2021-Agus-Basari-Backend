package category

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/database"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/helper"
	"github.com/labstack/echo/v4"
)

type CategoryRoutes struct{}

func (r CategoryRoutes) Route() []helper.Route {
	db := database.GetDBinstance()
	db.AutoMigrate(&Category{})
	catRepo := NewRepository(db)
	catServ := NewServices(catRepo)
	catHandler := NewHandler(catServ)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/categories",
			Handler: catHandler.AddCategory,
		},
		{
			Method:  echo.GET,
			Path:    "/categories",
			Handler: catHandler.ShowAllCategories,
		},
	}
}
