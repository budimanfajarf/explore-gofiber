package main

import (
	database "explore-gofiber/config"
	env "explore-gofiber/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	envConfig := env.LoadEnv()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + envConfig.ProjectPort)
}
