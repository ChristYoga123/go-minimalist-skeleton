package configs

import (
	"log"

	"gorm.io/gorm"
)

// RunMigration automatically migrates the database schema.
// It takes the GORM DB instance as a parameter.
func RunMigration(db *gorm.DB) {
	// --- PLACE TO REGISTER YOUR MODELS FOR MIGRATION ---
	// Add all your GORM models here to be auto-migrated.
	err := db.AutoMigrate(
	// &models.Product{}, // Example of another model
	// &models.Order{},   // Example of another model
	)

	if err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	log.Println("Database migrations completed successfully.")
}
