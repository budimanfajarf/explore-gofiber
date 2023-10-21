package http

import (
	"explore-gofiber/constant"

	"github.com/gofiber/fiber/v2"
)

func CustomHttpException(ctx *fiber.Ctx, status int, code string, message ...string) error {
	errorMessage := constant.DefaultErrMessage[code]

	if len(message) > 0 {
		errorMessage = message[0]
	}

	return ctx.Status(status).JSON(HttpResponse{
		Error: &HttpError{
			Code:    code,
			Message: errorMessage,
		},
	})
}

func InvalidCredentialsException(ctx *fiber.Ctx, message ...string) error {
	return CustomHttpException(ctx, fiber.StatusBadRequest, constant.ErrInvalidCredentials, message...)
}
