package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(err.Error())
}

func InternalError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
}

func SuccessResponse(ctx *fiber.Ctx, msg string, data any) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": msg,
		"data": data,
	})
}