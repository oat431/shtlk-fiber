package controller

import (
	"oat431/shtlk-fiber/common"
	"oat431/shtlk-fiber/payload/response"
	"oat431/shtlk-fiber/service"

	"github.com/gofiber/fiber/v3"
)

type ShortLinkController struct {
	service service.ShortLinkService
}

func NewShortLinkController(service service.ShortLinkService) *ShortLinkController {
	return &ShortLinkController{service: service}
}

func (s *ShortLinkController) GetAllShortLinks(c fiber.Ctx) error {
	shortLinkDTOs, err := s.service.GetAllLinks(c.Context())
	var res = common.ResponseDTO[[]response.ShortLinkDTO]{}
	if err != nil {
		res.Data = nil
		res.Status = common.ERROR
		res.Error = &common.ResponseDTOError{
			HttpCode:  fiber.StatusInternalServerError,
			ErrorCode: "INTERNAL_SERVER_ERROR",
			Message:   "Failed to retrieve short links",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	res.Data = &shortLinkDTOs
	res.Status = common.SUCCESS
	res.Error = nil

	return c.Status(fiber.StatusOK).JSON(res)
}
