package routes

import (
	"github.com/gofiber/fiber/v2"
	"user-service/app/http/controllers"
)

func SetupApiRoutes(app *fiber.App) {
	// Create a new group for API routes
	api := app.Group("/api")

	// Health check route
	api.Get("/health", controllers.HealthController().Check)
}
