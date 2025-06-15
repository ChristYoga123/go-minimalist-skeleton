//go:build wireinject
// +build wireinject

package wire

import (
	"go-skeleton/app/http/controllers"
	"go-skeleton/app/http/controllers/interfaces"
	"go-skeleton/app/repositories"
	"go-skeleton/app/services"

	"github.com/google/wire"
)

// Provider sets for each feature
// Add new provider set when adding new feature
// Example:
// var userSet = wire.NewSet(
//     repositories.NewUserRepository,
//     services.NewUserService,
//     controllers.NewUserController,
// )
var healthSet = wire.NewSet(
	repositories.NewHealthRepository,
	services.NewHealthService,
	controllers.NewHealthController,
)

// InitializeControllers initializes all controllers with their dependencies
// To add new feature:
// 1. Create new provider set above (e.g., userSet)
// 2. Add new controller interface to return values
// 3. Add new provider set to wire.Build
// 4. Add new controller to return statement
// Example:
// func InitializeControllers() (
//     interfaces.HealthController,
//     interfaces.UserController,  // New controller
//     error,
// ) {
//     wire.Build(
//         healthSet,
//         userSet,  // New provider set
//     )
//     return nil, nil, nil
// }
func InitializeControllers() (interfaces.HealthController, error) {
	wire.Build(healthSet)
	return nil, nil
}

