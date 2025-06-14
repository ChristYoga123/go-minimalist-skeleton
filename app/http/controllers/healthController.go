package controllers

import (
	"go-skeleton/app/services"
	"go-skeleton/utils"

	"github.com/gofiber/fiber/v2"
)

// HealthControllerStruct handles incoming HTTP requests for the health check endpoint.
type HealthControllerStruct struct {
	service services.HealthServiceInterface
}

// HealthController is the constructor for HealthController.
// It initializes its own dependency (HealthService), making it self-contained.
func HealthController() *HealthControllerStruct {
	return &HealthControllerStruct{
		service: services.HealthService(),
	}
}

// Check is the handler function for the GET /api/health route.
// It orchestrates the call to the service and formats the final JSON response.
func (h *HealthControllerStruct) Check(c *fiber.Ctx) error {
	dbStatus, err := h.service.CheckHealth()

	// Prepare the data payload for the response.
	healthData := map[string]string{
		"service_status":  "Running",
		"database_status": dbStatus,
	}

	// If the service returns an error, send a standardized error response.
	if err != nil {
		// Even in an error state, we provide context in the data payload.
		healthData["database_status"] = "Not Connected"
		return utils.ErrorResponse(
			c,
			fiber.StatusServiceUnavailable, // HTTP 503
			"One or more services are unavailable",
			healthData,
		)
	}

	// If everything is healthy, send a standardized success response.
	return utils.SuccessResponse(
		c,
		fiber.StatusOK, // HTTP 200
		"All services are healthy",
		healthData,
	)
}
