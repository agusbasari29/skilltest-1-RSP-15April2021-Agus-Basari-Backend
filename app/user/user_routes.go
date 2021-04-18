package user

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/auth"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/database"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/helper"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct{}

func (r UserRoutes) Route() []helper.Route {
	db := database.GetDBinstance()
	db.AutoMigrate(User{})
	repo := NewRepository(db)
	services := NewServices(repo)
	authService := auth.NewAuthService()
	handler := NewHandler(services, authService)

	return []helper.Route{

		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: handler.UserRegistration,
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: handler.UserLogin,
		}, {
			Method:  echo.GET,
			Path:    "/secret",
			Handler: handler.SecretResource,
		},
	}
}
