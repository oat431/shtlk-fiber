package controller

import (
	"oat431/shtlk-fiber/common"
	"oat431/shtlk-fiber/payload/request"
	"oat431/shtlk-fiber/payload/response"
	"oat431/shtlk-fiber/service"
	"oat431/shtlk-fiber/validate"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
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

func (s *ShortLinkController) CreateRandomShortLink(c fiber.Ctx) error {
	var req request.ShortLinkRequest
	err := validate.ValidateShortLinkRequest(req, c)
	if err != nil {
		log.Error("Validation error: ", err)
		return err
	}

	shortLinkDTO, err := s.service.CreateRandomShortLink(c.Context(), req.Url)
	var res = common.ResponseDTO[response.ShortLinkDTO]{}
	if err != nil {
		res.Data = nil
		res.Status = common.ERROR
		res.Error = &common.ResponseDTOError{
			HttpCode:  fiber.StatusInternalServerError,
			ErrorCode: "INTERNAL_SERVER_ERROR",
			Message:   "Failed to create short link",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	res.Data = shortLinkDTO
	res.Status = common.SUCCESS
	res.Error = nil

	return c.Status(fiber.StatusCreated).JSON(res)
}
