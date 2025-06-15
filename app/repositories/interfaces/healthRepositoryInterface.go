package interfaces

// HealthRepository defines the interface for health check repository
type HealthRepository interface {
	Check() error
} 