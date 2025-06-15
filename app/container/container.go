package container

import (
	containerInterfaces "go-skeleton/app/container/interfaces"
	"go-skeleton/app/http/controllers"
	controllerInterfaces "go-skeleton/app/http/controllers/interfaces"
	"go-skeleton/app/repositories"
	"go-skeleton/app/services"
)

// Container holds all dependencies
type Container struct {
	// Controllers
	healthController controllerInterfaces.HealthController
	// Add other controllers here
}

// NewContainer creates a new container with all dependencies
func NewContainer() containerInterfaces.Container {
	// Initialize repositories
	healthRepo := repositories.NewHealthRepository()

	// Initialize services
	healthService := services.NewHealthService(healthRepo)

	// Initialize controllers
	healthController := controllers.NewHealthController(healthService)

	return &Container{
		healthController: healthController,
		// Add other controllers here
	}
}

// GetHealthController returns the health controller
func (c *Container) GetHealthController() controllerInterfaces.HealthController {
	return c.healthController
} 