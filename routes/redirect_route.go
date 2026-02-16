package routes

import (
	"oat431/shtlk-fiber/controller"

	"github.com/gofiber/fiber/v3"
)

func RegisterRedirectRoutes(router fiber.Router, controller *controller.RedirectController) {
	random := router.Group("/r")
	custom := router.Group("/c")

	random.Get("/:code", controller.ShortLinkRedirect)
	custom.Get("/:code", controller.ShortLinkRedirect)

}
