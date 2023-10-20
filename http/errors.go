package http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

/*
 * Copied from helper, to prevent import cycle not allowed
 */
func snakeCaseToWords(input string) string {
	words := strings.Split(input, "_")
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, " ")
}

func HttpException(ctx *fiber.Ctx, status int, message ...string) error {
	errorCode := HttpCode[status]
	// errorMessage := utils.SnakeCaseToWords(errorCode)
	errorMessage := snakeCaseToWords(errorCode)

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
