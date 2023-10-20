package config

import (
	"errors"
	"explore-gofiber/http"

	"github.com/gofiber/fiber/v2"
)

// @see https://docs.gofiber.io/guide/error-handling/
func FiberConfig() fiber.Config {
	return fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			// err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			err = http.HttpException(ctx, code, err.Error())
			if err != nil {
				// In case the SendFile fails
				// return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
				return http.InternalServerErrorException(ctx)
			}

			// Return from handler
			return nil
		},
	}
}
