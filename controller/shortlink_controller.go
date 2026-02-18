package controller

import (
	"oat431/shtlk-fiber/common"
	"oat431/shtlk-fiber/payload/request"
	"oat431/shtlk-fiber/payload/response"
	"oat431/shtlk-fiber/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

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
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ResponseDTO[any]{
			Data:   nil,
			Status: common.ERROR,
			Error: &common.ResponseDTOError{
				HttpCode:  fiber.StatusBadRequest,
				ErrorCode: "BAD_REQUEST",
				Message:   "Invalid request body",
			},
		})
	}
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.ResponseDTO[any]{
			Data:   nil,
			Status: common.ERROR,
			Error: &common.ResponseDTOError{
				HttpCode:  fiber.StatusBadRequest,
				ErrorCode: "VALIDATION_ERROR",
				Message:   err.Error(),
			},
		})
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
