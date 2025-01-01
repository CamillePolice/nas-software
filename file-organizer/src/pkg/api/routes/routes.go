package routes

import (
	"file-organizer/src/pkg/api/handlers"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo) {
	// Initialize handlers
	h := handlers.NewOrganizerHandler()

	// Health group
	health := e.Group("/health")
	{
		health.GET("", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"status": "healthy",
			})
		})
	}

	// Binary group
	binary := e.Group("/binary")
	{
		binary.GET("/windows", h.GetWindowsBinary)
		binary.GET("/linux", h.GetLinuxBinary)
		binary.GET("/mac", h.GetMacBinary)
	}
}
