package interfaces

import "github.com/gofiber/fiber/v2"

// HealthController defines the interface for health check controller
type HealthController interface {
	Check(ctx *fiber.Ctx) error
} 