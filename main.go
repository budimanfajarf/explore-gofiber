package main

import (
	env "explore-gofiber/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	envConfig := env.LoadEnv()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + envConfig.ProjectPort)
}
