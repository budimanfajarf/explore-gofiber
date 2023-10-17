package main

import (
	"explore-gofiber/config"
	"explore-gofiber/database"
	"explore-gofiber/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.LoadEnv()
	database.Connect(envConfig)

	app := fiber.New()
	router.SetUpRoutes(app)
	app.Listen(":" + envConfig.ProjectPort)
}
