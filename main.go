package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"user-service/configs"
	"user-service/routes"
	_ "user-service/utils"
)

func main() {
	//Init Database
	if err := configs.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err.Error())
	}
	//Init Gofiber
	app := fiber.New()
	app.Use(logger.New())

	//Init Routes
	routes.SetupApiRoutes(app)

	//Start the server
	if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err.Error())
	}
}
