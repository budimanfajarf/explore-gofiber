package router

import (
	"explore-gofiber/http"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return http.Success(c, 200, "Hello World")
	})

	setUpRoutesV1(app)

	app.Use(func(c *fiber.Ctx) error {
		return http.NotFoundException(c)
	})
}
