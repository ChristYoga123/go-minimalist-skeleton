package repositories

import (
	"go-skeleton/app/repositories/interfaces"
	"go-skeleton/configs"

	"gorm.io/gorm"
)

type healthRepository struct {
	db *gorm.DB
}

// NewHealthRepository creates a new instance of health repository
func NewHealthRepository() interfaces.HealthRepository {
	return &healthRepository{
		db: configs.GetDB(),
	}
}

// Check implements the health check repository method
func (r *healthRepository) Check() error {
	// Check database connection
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}

	// Ping database
	return sqlDB.Ping()
} 