package routes

import (
	"oat431/shtlk-fiber/internal/dto/request"
	"oat431/shtlk-fiber/internal/middleware"

	"github.com/gofiber/fiber/v3"

	"oat431/shtlk-fiber/internal/controller"
)

func RegisterShortLinkRoutes(router fiber.Router, controller *controller.ShortLinkController) {
	route := router.Group("/short-link")

	route.Get("/", controller.GetAllShortLinks)
	route.Post("/random",
		middleware.Validate[request.ShortLinkRequest],
		controller.CreateRandomShortLink,
	)
	route.Post("/custom",
		middleware.Validate[request.ShortLinkRequest],
		controller.CreateCustomShortLink,
	)
}
