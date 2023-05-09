package routes

import (
	"prakerja3/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.InsertUserController)
	return e
}
