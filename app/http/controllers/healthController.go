package controllers

import (
	"go-skeleton/configs"
	"go-skeleton/utils"

	"github.com/gofiber/fiber/v2"
)

// HealthController handles health check endpoints
type HealthController struct{}

// Check is the handler function for the GET /api/health route
func (h *HealthController) Check(c *fiber.Ctx) error {
	// Get database instance
	db := configs.GetDB()
	
	// Check database connection
	sqlDB, err := db.DB()
	if err != nil {
		return utils.ErrorResponse(
			c,
			fiber.StatusServiceUnavailable,
			"Database connection error",
			map[string]string{
				"service_status":  "Running",
				"database_status": "Not Connected",
			},
		)
	}

	// Ping the database
	if err := sqlDB.Ping(); err != nil {
		return utils.ErrorResponse(
			c,
			fiber.StatusServiceUnavailable,
			"Database ping failed",
			map[string]string{
				"service_status":  "Running",
				"database_status": "Not Connected",
			},
		)
	}

	// If everything is healthy, send success response
	return utils.SuccessResponse(
		c,
		fiber.StatusOK,
		"All services are healthy",
		map[string]string{
			"service_status":  "Running",
			"database_status": "Connected",
		},
	)
}
