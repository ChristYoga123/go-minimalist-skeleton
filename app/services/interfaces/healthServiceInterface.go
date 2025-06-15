package interfaces

// HealthService defines the interface for health check service
type HealthService interface {
	Check() error
} 