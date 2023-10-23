package http

import (
	"strings"
	"unicode"

	"github.com/gofiber/fiber/v2"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
)

func Success(ctx *fiber.Ctx, status int, data interface{}) error {
	return ctx.Status(status).JSON(HttpResponse{
		Data: data,
	})
}

func SuccessWithMeta(ctx *fiber.Ctx, status int, data interface{}, meta interface{}) error {
	return ctx.Status(status).JSON(HttpResponse{
		Data: data,
		Meta: meta,
	})
}

func StatusCodeStr(status int) string {
	message := fiberUtils.StatusMessage(status)
	words := strings.Fields(message)
	var result strings.Builder

	for i, word := range words {
		if i > 0 {
			result.WriteRune('_')
		}

		for _, char := range word {
			if !unicode.IsLetter(char) {
				char = '_'
			}
			result.WriteRune(unicode.ToUpper(char))
		}
	}

	return result.String()
}

func Exception(ctx *fiber.Ctx, status int, message ...string) error {
	errorCode := StatusCodeStr(status)
	errorMessage := fiberUtils.StatusMessage(status)

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
