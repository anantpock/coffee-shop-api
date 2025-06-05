package main

import (
	"github.com/anantpock/coffee-shop-api/config"
	"github.com/anantpock/coffee-shop-api/models"
	"github.com/anantpock/coffee-shop-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
