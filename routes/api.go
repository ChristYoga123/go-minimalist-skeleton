package routes

import (
	"go-skeleton/app/wire"

	"github.com/gofiber/fiber/v2"
)

// SetupApiRoutes sets up all API routes
// To add new feature:
// 1. Add new controller to wire.InitializeControllers() return values
// 2. Create new route group for the feature
// 3. Add routes using the new controller
// Example:
// func SetupApiRoutes(app *fiber.App) {
//     api := app.Group("/api")
//
//     // Initialize controllers
//     healthController, userController, err := wire.InitializeControllers()
//     if err != nil {
//         panic(err)
//     }
//
//     // Health routes
//     api.Get("/health", healthController.Check)
//
//     // User routes
//     userGroup := api.Group("/users")
//     userGroup.Post("/", userController.Create)
//     userGroup.Get("/:id", userController.GetByID)
//     userGroup.Get("/", userController.GetAll)
//     userGroup.Put("/:id", userController.Update)
//     userGroup.Delete("/:id", userController.Delete)
// }
func SetupApiRoutes(app *fiber.App) {
	// Create a new group for API routes
	api := app.Group("/api")

	// Initialize dependencies using wire
	healthController, err := wire.InitializeControllers()
	if err != nil {
		panic(err)
	}

	// Health check route
	api.Get("/health", healthController.Check)
}
