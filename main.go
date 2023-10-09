package main

import (
	database "explore-gofiber/config"
	env "explore-gofiber/config"
	"explore-gofiber/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	envConfig := env.LoadEnv()

	router.SetUpRoutes(app)

	app.Listen(":" + envConfig.ProjectPort)
}
