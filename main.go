package main

import (
	"explore-gofiber/config"
	"explore-gofiber/database"
	"explore-gofiber/modules"

	"github.com/gofiber/fiber/v2"
)

func main() {
	env := config.LoadEnv()
	database.Connect()
	modules.Init()

	app := fiber.New()
	setUpRoutes(app)

	app.Listen(":" + env.ProjectPort)
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// --- api v1 ---
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	v1.Get("/articles", modules.ArticleHandler.GetList)
	// --- api v1 ---

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

}
