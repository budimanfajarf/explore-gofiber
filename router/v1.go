package router

import (
	"explore-gofiber/article"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesV1(app *fiber.App) {
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		// c.JSON(fiber.Map{
		// 	"message": "🐣 v1",
		// })
		return c.Next()
	})

	v1.Get("/articles", article.GetArticles)
}
