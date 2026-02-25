package routes

import (
	"oat431/shtlk-fiber/internal/bootstrap"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func init() {
	log.Info("Initializing routes...")
}

func SetupRoutes(app *fiber.App, container *bootstrap.AppContainer) {
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
	RegisterRedirectRoutes(app, container.RedirectController)
}
