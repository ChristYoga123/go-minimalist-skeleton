package interfaces

import "context"

// CommonRepository defines common CRUD operations for repositories
type CommonRepository[T any] interface {
	// Create creates a new entity
	Create(ctx context.Context, entity *T) error

	// FindByID finds an entity by its ID
	FindByID(ctx context.Context, id uint) (*T, error)

	// FindAll retrieves all entities with optional pagination
	FindAll(ctx context.Context, page, limit int) ([]*T, error)

	// Update updates an existing entity
	Update(ctx context.Context, entity *T) error

	// Delete removes an entity by its ID
	Delete(ctx context.Context, id uint) error

	// Count returns the total number of entities
	Count(ctx context.Context) (int64, error)
} 