package controllers

import (
	"net/http"

	"github.com/anantpock/coffee-shop-api/config"
	"github.com/anantpock/coffee-shop-api/models"
	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, order.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	order.Total = float64(order.Quantity) * product.Price

	config.DB.Create(&order)
	c.JSON(http.StatusOK, order)
}
