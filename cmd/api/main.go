package main

import (
	"oat431/shtlk-fiber/internal/bootstrap"
	"oat431/shtlk-fiber/internal/config"
	"oat431/shtlk-fiber/internal/routes"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	config.LoadEnvConfig()

	// Initialize database connection
	db := config.StartDatabase()
	defer db.Close()

	// Initialize application container with dependencies
	container := bootstrap.NewAppContainer(db)

	// Create Fiber app
	app := fiber.New()

	// Setup Routes
	routes.SetupRoutes(app, container)

	port := os.Getenv("PORT")
	// Start the server
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("port 8000 is already in use")
	}
}
