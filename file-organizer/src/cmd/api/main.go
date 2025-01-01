// cmd/api/main.go
package main

import (
	"file-organizer/src/pkg/api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

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
