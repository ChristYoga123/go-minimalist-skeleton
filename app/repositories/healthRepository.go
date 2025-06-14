package repositories

import (
	"go-skeleton/configs"

	"gorm.io/gorm"
)

// HealthRepositoryInterface defines the interface for health check operations.
type HealthRepositoryInterface interface {
	CheckDatabase() error
}

// healthRepository implements the HealthRepository interface.
// It holds a reference to the GORM database connection.
type healthRepositoryStruct struct {
	db *gorm.DB
}

// HealthRepository is the constructor for healthRepository.
// It retrieves the singleton database instance from the configs package,
// so no parameters are needed.
func HealthRepository() HealthRepositoryInterface {
	return &healthRepositoryStruct{db: configs.GetDB()}
}

// CheckDatabase performs a ping to the database to verify the connection.
func (r *healthRepositoryStruct) CheckDatabase() error {
	// Get the underlying sql.DB object from GORM
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	// Ping the database
	return sqlDB.Ping()
}
