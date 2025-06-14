package utils

import "github.com/gofiber/fiber/v2"

// ResponseFormatter defines the standard structure for API responses.
type ResponseFormatter struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Use omitempty to hide null data fields
}

// SuccessResponse creates a standardized success response.
// It takes a status code, a message, and the data payload.
func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(ResponseFormatter{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// ErrorResponse creates a standardized error response.
// It takes a status code, a message, and optional error details (can be nil).
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, errorDetails interface{}) error {
	return c.Status(statusCode).JSON(ResponseFormatter{
		Status:  "error",
		Message: message,
		Data:    errorDetails, // This can hold validation errors or other info
	})
}
