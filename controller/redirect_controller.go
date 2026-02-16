package controller

import (
	"oat431/shtlk-fiber/service"

	"github.com/gofiber/fiber/v3"
)

type RedirectController struct {
	service.ShortLinkService
}

func NewRedirectController(shortLinkService service.ShortLinkService) *RedirectController {
	return &RedirectController{ShortLinkService: shortLinkService}
}

func (s *RedirectController) ShortLinkRedirect(c fiber.Ctx) error {
	return c.Status(fiber.StatusMovedPermanently).Redirect().To("https://github.com/oat431")
}
