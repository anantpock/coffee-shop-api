package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/anantpock/coffee-shop-api/config"
	"github.com/anantpock/coffee-shop-api/models"
	"github.com/anantpock/coffee-shop-api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	input.Password = string(hashedPassword)
	if input.Role == "" {
		input.Role = "customer"
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role, os.Getenv("JWT_SECRET"), time.Hour*24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
