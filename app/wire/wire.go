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

var healthSet = wire.NewSet(
	repositories.NewHealthRepository,
	services.NewHealthService,
	controllers.NewHealthController,
)

// Add new interfaces to return value as needed
func InitializeControllers() (interfaces.HealthController, error) {
	wire.Build(healthSet)
	// Add any additional sets or providers here if needed
	return nil, nil
}
