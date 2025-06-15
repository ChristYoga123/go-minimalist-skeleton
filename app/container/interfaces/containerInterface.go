package interfaces

import (
	"go-skeleton/app/http/controllers/interfaces"
)

// Container defines the interface for dependency container
type Container interface {
	// Controllers
	GetHealthController() interfaces.HealthController
	// Add other getters here
} 