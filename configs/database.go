package configs

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// DBConfig holds all configuration for database connections, parsed from environment variables.
type DBConfig struct {
	DBDriver   string `env:"DB_DRIVER,required"`
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBDatabase string `env:"DB_DATABASE"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
}

// loadDBConfig parses environment variables into the DBConfig struct.
func loadDBConfig() (*DBConfig, error) {
	cfg := &DBConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}
	return cfg, nil
}

// NewDatabaseConnection is a factory function that creates a database connection
// based on the driver specified in the configuration.
// It uses a switch case to handle different SQL and NoSQL databases.
// To add a new database, simply add a new case to this switch.
func NewDatabaseConnection(config *DBConfig) (db *gorm.DB, err error) {
	var dialector gorm.Dialector

	switch config.DBDriver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBDatabase)
		dialector = mysql.Open(dsn)
		log.Println("Connecting to MySQL...")
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			config.DBHost, config.DBUsername, config.DBPassword, config.DBDatabase, config.DBPort)
		dialector = postgres.Open(dsn)
		log.Println("Connecting to PostgreSQL...")
	case "sqlite":
		// For sqlite, DB_DATABASE is the path to the .db file. Example: "gorm.db"
		dsn := config.DBDatabase
		dialector = sqlite.Open(dsn)
		log.Println("Connecting to SQLite...")
	case "sqlserver":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBDatabase)
		dialector = sqlserver.Open(dsn)
	//	log.Println("Connecting to SQL Server...")
	// --- PLACE TO REGISTER NEW NOSQL DATABASES ---
	case "redis":
		// Redis connection logic here
		log.Println("Connecting to Redis...")
		// return redisClient, nil
	case "mongo":
		// MongoDB connection logic here
		log.Println("Connecting to MongoDB...")
		// return mongoClient, nil
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", config.DBDriver)
	}

	// Establish GORM connection for SQL databases
	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
