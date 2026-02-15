package routes

import (
	"github.com/gofiber/fiber/v3"

	"oat431/shtlk-fiber/controller"
)

func RegisterShortLinkRoutes(router fiber.Router, controller *controller.ShortLinkController) {
	route := router.Group("/short-link")

	route.Get("/", controller.GetAllShortLinks)
}
