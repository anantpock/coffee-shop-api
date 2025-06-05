package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/anantpock/coffee-shop-api/utils"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header missing"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseToken(tokenStr, os.Getenv("JWT_SECRET"))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		// Save values to context so handlers can use them
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		// Continue processing next middleware/handler
		return next(c)
	}
}
