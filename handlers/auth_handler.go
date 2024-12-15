package handlers

import (
	"gorm_crud/auth"
	"gorm_crud/config"
	"gorm_crud/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// Mock user authentication (replace with DB logic)
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	// Example: Check if user exists in the database
	var dbUser model.User
	if err := config.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "user not found"})
	}

	// Generate a token
	token, err := auth.GenerateToken(dbUser.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
