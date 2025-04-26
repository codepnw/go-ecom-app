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

func BadRequestResponse(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
		"message": msg,
	})
}

func NotFoundResponse(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
		"message": msg,
	})
}

func SuccessResponse(ctx *fiber.Ctx, msg string, data any) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": msg,
		"data": data,
	})
}

func SuccessCreated(ctx *fiber.Ctx, msg string, data any) error {
	return ctx.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": msg,
		"data": data,
	})
}

func NoContentResponse(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusNoContent).JSON(nil)
}