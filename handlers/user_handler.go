package handlers

import (
	"gorm_crud/config"
	"gorm_crud/model"
	"gorm_crud/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateUser creates a new user.
func CreateUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid payload", nil)
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Failed to create user", nil)
	}

	return utils.SendResponse(c, http.StatusCreated, "User created successfully", user)
}

// GetUser retrieves a user by ID.
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
	}

	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return utils.SendResponse(c, http.StatusNotFound, "User not found", nil)
	}

	return utils.SendResponse(c, http.StatusOK, "User retrieved successfully", user)
}

// UpdateUser updates a user's details.
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
	}

	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return utils.SendResponse(c, http.StatusNotFound, "User not found", nil)
	}

	if err := c.Bind(&user); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid payload", nil)
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Failed to update user", nil)
	}

	return utils.SendResponse(c, http.StatusOK, "User updated successfully", user)
}

// DeleteUser deletes a user by ID.
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid user ID", nil)
	}

	if err := config.DB.Delete(&model.User{}, id).Error; err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, "Failed to delete user", nil)
	}

	return utils.SendResponse(c, http.StatusOK, "User deleted successfully", nil)
}
