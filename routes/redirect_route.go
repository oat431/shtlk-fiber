package routes

import (
	"oat431/shtlk-fiber/controller"

	"github.com/gofiber/fiber/v3"
)

func RegisterRedirectRoutes(router fiber.Router, controller *controller.RedirectController) {
	router.Get("/:linkType/:code", controller.ShortLinkRedirect)
}
