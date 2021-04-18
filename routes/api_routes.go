package routes

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/book"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/category"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/helper"
	"github.com/labstack/echo/v4"
)

func DefineApiRoutes(e *echo.Echo) {
	handlers := []helper.Handler{
		category.CategoryRoutes{},
		book.BookRoutes{},
	}

	var routes []helper.Route

	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}

	api := e.Group("/api")

	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
			}
		}
	}
}
