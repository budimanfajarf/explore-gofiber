package http

import (
	"github.com/gofiber/fiber/v2"
)

func HttpException(ctx *fiber.Ctx, status int, defaultMessage string, message ...string) error {
	errorMessage := defaultMessage

	if len(message) > 0 {
		errorMessage = message[0]
	}

	return ctx.Status(status).JSON(HttpResponse{
		Error: &HttpError{
			Code:    status,
			Message: errorMessage,
		},
	})
}

func InternalServerErrorException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusInternalServerError, "Internal Server Error", message...)
}

func BadRequestException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusBadRequest, "Bad Request", message...)
}

func UnauthorizedException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusUnauthorized, "Unauthorized", message...)
}

func ForbiddenException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusForbidden, "Forbidden", message...)
}

func NotFoundException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusNotFound, "Not Found", message...)
}
