package routes

import (
	"go-skeleton/app/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupApiRoutes(app *fiber.App) {
	// Create a new group for API routes
	api := app.Group("/api")

	// Health check route
	api.Get("/health", (&controllers.HealthController{}).Check)
}
