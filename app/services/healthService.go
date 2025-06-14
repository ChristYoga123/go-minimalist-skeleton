package services

import "user-service/app/repositories"

// HealthServiceInterface defines the interface for health check business logic.
// This decouples the controller from the concrete service implementation.
type HealthServiceInterface interface {
	CheckHealth() (dbStatus string, err error)
}

// healthService is the concrete implementation of HealthService.
// It depends on an implementation of HealthRepository.
type healthServiceStruct struct {
	repo repositories.HealthRepositoryInterface
}

// HealthService is the constructor for healthService.
// It creates its own dependency (HealthRepository) automatically.
// This simplifies initialization in the routing layer.
func HealthService() HealthServiceInterface {
	return &healthServiceStruct{
		repo: repositories.HealthRepository(),
	}
}

// CheckHealth contains the business logic for the health check.
// It translates the result from the repository into a business-friendly status.
func (s *healthServiceStruct) CheckHealth() (string, error) {
	err := s.repo.CheckDatabase()
	if err != nil {
		return "Not Connected", err
	}
	return "Connected", nil
}
