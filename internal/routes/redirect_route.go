package routes

import (
	"oat431/shtlk-fiber/internal/controller"

	"github.com/gofiber/fiber/v3"
)

func RegisterRedirectRoutes(router fiber.Router, controller *controller.RedirectController) {
	router.Get("/:linkType/:code", controller.ShortLinkRedirect)
}
