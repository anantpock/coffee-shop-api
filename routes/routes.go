package routes

import (
	"github.com/anantpock/coffee-shop-api/controllers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "pong"})
	})

	e.GET("/users", controllers.GetUsers)
	e.POST("/users", controllers.CreateUser)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Welcome to the Coffee Shop API ☕",
		})
	})
}
