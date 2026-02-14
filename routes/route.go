package routes

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func init() {
	log.Info("Initializing routes...")
}

func StartingApplication() {
	app := fiber.New()
	port := os.Getenv("PORT")

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("port 8000 is already in use")
	}
}
