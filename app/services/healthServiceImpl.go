package services

import (
	repoInterfaces "go-skeleton/app/repositories/interfaces"
	serviceInterfaces "go-skeleton/app/services/interfaces"
)

type healthService struct {
	healthRepo repoInterfaces.HealthRepository
}

// NewHealthService creates a new instance of health service
func NewHealthService(healthRepo repoInterfaces.HealthRepository) serviceInterfaces.HealthService {
	return &healthService{
		healthRepo: healthRepo,
	}
}

// Check implements the health check service method
func (s *healthService) Check() error {
	return s.healthRepo.Check()
} 