package routes

import (
	"oat431/shtlk-fiber/bootstrap"
	"oat431/shtlk-fiber/config"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func init() {
	log.Info("Initializing routes...")
}

func StartingApplication() {
	// Initialize database connection
	db := config.StartDatabase()
	defer db.Close()

	// Initialize application container with dependencies
	container := bootstrap.NewAppContainer(db)

	// Create Fiber app
	app := fiber.New()
	port := os.Getenv("PORT")

	// Middleware to log incoming requests
	app.Use(func(c fiber.Ctx) error {
		log.Info("", c.Method(), " ", c.Path())
		err := c.Next()
		if err != nil {
			log.Error("Error occurred while processing request: ", err)
		}
		return err
	})

	// Register routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Register routes for different modules
	RegisterHealthRoutes(v1)
	RegisterShortLinkRoutes(v1, container.ShortLinkController)

	// Start the server
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("port 8000 is already in use")
	}
}
