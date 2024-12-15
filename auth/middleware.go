package auth

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware validates the JWT token.
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing authorization header"})
		}

		// Extract token
		tokenString := authHeader[len("Bearer "):]
		token, err := ValidateToken(tokenString)
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid or expired token"})
		}

		// Set user ID in context (optional)
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("user_id", claims["user_id"])
		}
		fmt.Println(c, "claims")

		return next(c)
	}
}
