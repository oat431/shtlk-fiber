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

	db := config.StartDatabase()
	defer db.Close()

	container := bootstrap.NewAppContainer(db)

	app := fiber.New()
	routes.SetupRoutes(app, container)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("port 8000 is already in use")
	}
}
