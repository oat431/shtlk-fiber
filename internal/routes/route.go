package routes

import (
	"oat431/shtlk-fiber/internal/bootstrap"
	"oat431/shtlk-fiber/internal/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func init() {
	log.Info("Initializing routes...")
}

func SetupRoutes(app *fiber.App, container *bootstrap.AppContainer) {
	app.Use(middleware.RequestMiddleware)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	RegisterHealthRoutes(v1)
	RegisterShortLinkRoutes(v1, container.ShortLinkController)
	RegisterRedirectRoutes(app, container.RedirectController)
}
