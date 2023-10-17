package main

import (
	"explore-gofiber/config"
	"explore-gofiber/database"
	"explore-gofiber/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	env := config.LoadEnv()
	database.Connect(env)
	app := fiber.New()
	router.SetUpRoutes(app)
	app.Listen(":" + env.ProjectPort)
}
