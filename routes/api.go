package routes

import (
	"go-skeleton/app/http/controllers"
	"go-skeleton/app/repositories"
	"go-skeleton/app/services"

	"github.com/gofiber/fiber/v2"
)

func SetupApiRoutes(app *fiber.App) {
	// Create a new group for API routes
	api := app.Group("/api")

	// Initialize dependencies
	healthRepo := repositories.NewHealthRepository()
	healthService := services.NewHealthService(healthRepo)
	healthController := controllers.NewHealthController(healthService)

	// Health check route
	api.Get("/health", healthController.Check)
}
