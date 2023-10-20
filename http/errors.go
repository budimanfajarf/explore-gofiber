package http

import "github.com/gofiber/fiber/v2"

func HttpException(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(HttpResponse{
		Error: &HttpError{
			Code:    status,
			Message: message,
		},
	})
}

func InternalServerErrorException(ctx *fiber.Ctx, message string) error {
	return HttpException(ctx, fiber.StatusInternalServerError, message)
}

func DefaultInternalServerErrorException(ctx *fiber.Ctx) error {
	return InternalServerErrorException(ctx, "Internal Server Error")
}

func BadRequestException(ctx *fiber.Ctx, message string) error {
	return HttpException(ctx, fiber.StatusBadRequest, message)
}

func UnauthorizedException(ctx *fiber.Ctx, message string) error {
	return HttpException(ctx, fiber.StatusUnauthorized, message)
}

func ForbiddenException(ctx *fiber.Ctx, message string) error {
	return HttpException(ctx, fiber.StatusForbidden, message)
}

func NotFoundException(ctx *fiber.Ctx, message string) error {
	return HttpException(ctx, fiber.StatusNotFound, message)
}

func DefaultNotFoundException(ctx *fiber.Ctx) error {
	return NotFoundException(ctx, "Not Found")
}
