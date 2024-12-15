package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm_crud/config"
	"gorm_crud/routes"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize database
	config.InitDB()

	// Register routes
	routes.RegisterRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
