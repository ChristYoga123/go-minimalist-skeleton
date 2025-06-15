package routes

import (
	"go-skeleton/app/container"

	"github.com/gofiber/fiber/v2"
)

func SetupApiRoutes(app *fiber.App) {
	// Create a new group for API routes
	api := app.Group("/api")

	// Initialize container
	container := container.NewContainer()

	// Health check route
	api.Get("/health", container.GetHealthController().Check)
}
