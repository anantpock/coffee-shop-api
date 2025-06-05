package controllers

import (
	"net/http"

	"github.com/anantpock/coffee-shop-api/config"
	"github.com/anantpock/coffee-shop-api/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}
