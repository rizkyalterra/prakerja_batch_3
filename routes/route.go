package routes

import (
	"prakerja3/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.InsertUserController)
	e.Use(middleware.Logger())
	return e
}
