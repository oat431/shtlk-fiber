package controller

import (
	"oat431/shtlk-fiber/common"
	"oat431/shtlk-fiber/payload/response"

	"github.com/gofiber/fiber/v3"
)

func GetInfo(c fiber.Ctx) error {
	info := response.InfoResponse{
		Name:        "Shortlink API",
		Description: "An API for creating and managing short links. Built with Go Fiber.",
		Version:     "1.0.0-dev",
	}

	res := common.ResponseDTO[response.InfoResponse]{
		Data:   info,
		Status: common.SUCCESS,
		Error:  nil,
	}

	return c.JSON(res)
}
