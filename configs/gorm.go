package configs

import (
	"log"
	"sync"

	"gorm.io/gorm"
)

var (
	// Singleton instances for database connections.
	// We use a pointer to gorm.DB so we can have a nil value when not connected.
	gormDB *gorm.DB

	// You can add other DB clients here, e.g., for Redis or MongoDB
	// redisClient *redis.Client
	// mongoClient *mongo.Client

	// sync.Once ensures that the initialization code is executed only once,
	// making the singleton pattern thread-safe.
	once sync.Once
)

// InitDatabase initializes the database connection using the factory function.
// It implements the Singleton pattern using sync.Once to ensure it's only called once.
func InitDatabase() error {
	var err error
	once.Do(func() {
		// Load configuration from environment variables
		cfg, loadErr := loadDBConfig()
		if loadErr != nil {
			err = loadErr
			return
		}

		// Create a new connection using the factory
		db, connErr := NewDatabaseConnection(cfg)
		if connErr != nil {
			err = connErr
			return
		}

		// --- ASSIGN TO SINGLETON INSTANCE ---
		// We assign the connection to our package-level variable.
		// If you were connecting to multiple DBs, you would check the type here.
		gormDB = db

		// Run migrations after a successful connection
		if gormDB != nil {
			log.Println("Database connection successful. Running migrations...")
			RunMigration(gormDB)
		}
	})

	return err
}

// GetDB returns the singleton GORM database instance.
// The rest of the application will use this function to get the DB connection.
func GetDB() *gorm.DB {
	if gormDB == nil {
		log.Fatal("Database is not initialized. Call InitDatabase() first.")
	}
	return gormDB
}

// Add other getters for NoSQL databases if needed
// func GetRedis() *redis.Client {
//     return redisClient
// }
