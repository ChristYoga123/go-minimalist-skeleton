package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// The init() function is a special Go function that runs automatically
// when this package is first imported. This is a perfect place to load
// environment variables from a .env file, making them available to the
// entire application.
func init() {
	// godotenv.Load() will load the .env file from the current directory
	// and set the environment variables. If it fails, it will print a
	// message but won't stop the application, as env vars might be set
	// by the system itself (e.g., in Docker or Kubernetes).
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from system environment variables")
	}
}
