package main

import (
	"explore-gofiber/config"
	"explore-gofiber/database"
	"explore-gofiber/http"
	"explore-gofiber/modules"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	env := config.LoadEnv()
	database.Connect()
	modules.Init()

	app := fiber.New(config.FiberConfig())
	// app.Use(recover.New()) // disable it to avoid confusion when getting runtime errors
	setUpRoutes(app)

	app.Listen(":" + env.ProjectPort)
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return http.Success(c, 200, "Hello World")
	})

	// --- api v1 ---
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	v1.Get("/articles", modules.ArticleHandler.GetList)
	v1.Get("/articles/:id", modules.ArticleHandler.GetDetails)
	// --- api v1 ---

	app.Use(func(c *fiber.Ctx) error {
		return http.NotFoundException(c)
	})

}
