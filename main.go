package main

import (
	"explore-gofiber/config"
	"explore-gofiber/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDb()
	app := fiber.New()

	envConfig := config.LoadEnv()

	router.SetUpRoutes(app)

	app.Listen(":" + envConfig.ProjectPort)
}
