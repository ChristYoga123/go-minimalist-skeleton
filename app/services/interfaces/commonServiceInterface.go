package interfaces

import "context"

// CommonService defines common CRUD operations for services
type CommonService[T any] interface {
	// Create creates a new entity
	Create(ctx context.Context, entity *T) error

	// GetByID retrieves an entity by its ID
	GetByID(ctx context.Context, id uint) (*T, error)

	// GetAll retrieves all entities with optional pagination
	GetAll(ctx context.Context, page, limit int) ([]*T, error)

	// Update updates an existing entity
	Update(ctx context.Context, entity *T) error

	// Delete removes an entity by its ID
	Delete(ctx context.Context, id uint) error

	// GetTotal returns the total number of entities
	GetTotal(ctx context.Context) (int64, error)
} 