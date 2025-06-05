package main

import (
	"log"

	"github.com/anantpock/coffee-shop-api/config"
	"github.com/anantpock/coffee-shop-api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected successfully")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
