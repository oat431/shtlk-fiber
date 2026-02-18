package validate

import (
	"oat431/shtlk-fiber/common"
	"oat431/shtlk-fiber/payload/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

func ValidateShortLinkRequest(req request.ShortLinkRequest, c fiber.Ctx) error {
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
	return nil
}
