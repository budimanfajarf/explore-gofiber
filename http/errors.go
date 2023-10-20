package http

import (
	"github.com/gofiber/fiber/v2"
)

func HttpException(ctx *fiber.Ctx, status int, message ...string) error {
	errorCode := HttpCode[status]
	errorMessage := errorCode

	if len(message) > 0 {
		errorMessage = message[0]
	}

	return ctx.Status(status).JSON(HttpResponse{
		Error: &HttpError{
			Code:    errorCode,
			Message: errorMessage,
		},
	})
}

func InternalServerErrorException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusInternalServerError, message...)
}

func BadRequestException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusBadRequest, message...)
}

func UnauthorizedException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusUnauthorized, message...)
}

func ForbiddenException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusForbidden, message...)
}

func NotFoundException(ctx *fiber.Ctx, message ...string) error {
	return HttpException(ctx, fiber.StatusNotFound, message...)
}
