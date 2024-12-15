package routes

import (
	"gorm_crud/auth"
	"gorm_crud/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/login", handlers.Login)

	// Protected routes
	protected := e.Group("/protected")
	protected.Use(auth.AuthMiddleware)

	protected.GET("/users/:id", handlers.GetUser)
	protected.POST("/users", handlers.CreateUser)
	protected.PUT("/users/:id", handlers.UpdateUser)
	protected.DELETE("/users/:id", handlers.DeleteUser)
}
