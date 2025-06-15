package controllers

import (
	controllerInterfaces "go-skeleton/app/http/controllers/interfaces"
	serviceInterfaces "go-skeleton/app/services/interfaces"
	"go-skeleton/utils"

	"github.com/gofiber/fiber/v2"
)

type healthController struct {
	healthService serviceInterfaces.HealthService
}

// NewHealthController creates a new instance of health controller
func NewHealthController(healthService serviceInterfaces.HealthService) controllerInterfaces.HealthController {
	return &healthController{
		healthService: healthService,
	}
}

// Check handles the health check endpoint
func (c *healthController) Check(ctx *fiber.Ctx) error {
	if err := c.healthService.Check(); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Service unhealthy", err)
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Service healthy", nil)
} 