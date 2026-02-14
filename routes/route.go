package routes

import (
	"oat431/shtlk-fiber/controller"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

func init() {
	log.Info("Initializing routes...")
}

func StartingApplication() {
	app := fiber.New()
	port := os.Getenv("PORT")

	api := app.Group("/api")
	v1 := api.Group("/v1")
	health := v1.Group("/health-check")

	health.Get("/health", healthcheck.New())
	health.Get("/info", controller.GetInfo)

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("port 8000 is already in use")
	}
}
