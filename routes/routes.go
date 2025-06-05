package routes

import (
	"github.com/anantpock/coffee-shop-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Coffee Shop API ☕",
		})
	})

}
