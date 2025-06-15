package repositories

import (
	"go-skeleton/app/repositories/interfaces"
)

type healthRepository struct{}

// NewHealthRepository creates a new instance of health repository
func NewHealthRepository() interfaces.HealthRepository {
	return &healthRepository{}
}

// Check implements the health check repository method
func (r *healthRepository) Check() error {
	// Add any database or external service checks here
	return nil
} 