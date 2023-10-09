package router

import (
	"explore-gofiber/article"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		// c.JSON(fiber.Map{
		// 	"message": "🐣 v1",
		// })
		return c.Next()
	})

	v1.Get("/articles", article.GetArticles)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
