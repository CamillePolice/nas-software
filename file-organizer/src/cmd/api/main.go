// cmd/api/main.go
package main

import (
	"file-organizer/src/pkg/api/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	err := godotenv.Load("../../../.env")
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Debug = true

	// Configure routes
	routes.ConfigureRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
