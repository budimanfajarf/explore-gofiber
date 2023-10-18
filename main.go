package main

import (
	"explore-gofiber/config"
	"explore-gofiber/database"
	"explore-gofiber/modules/article"

	"github.com/gofiber/fiber/v2"
)

func main() {
	env := config.LoadEnv()
	database.Connect()
	db := database.MySQL

	app := fiber.New()

	articleRepository := article.NewArticleRepository(db)
	articleService := article.NewArticleService(articleRepository)

	articleHandler := article.NewArticleHandler(articleService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	v1.Get("/articles", articleHandler.GetList)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":" + env.ProjectPort)
}
