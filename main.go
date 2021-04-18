package main

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/category"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/database"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/routes"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	db      = database.GetDBinstance()
	catRepo = category.NewRepository(db)
	catSer  = category.NewServices(catRepo)
	catHand = category.NewHandler(catSer)
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// e.POST("/books", catHand.AddCategory)
	e.Pre(middleware.RemoveTrailingSlash())
	routes.DefineApiRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
