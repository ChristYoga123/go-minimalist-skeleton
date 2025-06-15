package interfaces

import "github.com/gofiber/fiber/v2"

// CommonController defines common CRUD operations for controllers
type CommonController interface {
	// Create handles the creation of a new entity
	Create(ctx *fiber.Ctx) error

	// GetByID handles retrieving an entity by its ID
	GetByID(ctx *fiber.Ctx) error

	// GetAll handles retrieving all entities with pagination
	GetAll(ctx *fiber.Ctx) error

	// Update handles updating an existing entity
	Update(ctx *fiber.Ctx) error

	// Delete handles removing an entity by its ID
	Delete(ctx *fiber.Ctx) error
} 