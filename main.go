package main

import (
	"explore-gofiber/config"
	"explore-gofiber/database"
	"explore-gofiber/modules"
	"explore-gofiber/router"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	env := config.LoadEnv()
	database.Connect()
	modules.Init()

	app := fiber.New(config.FiberConfig())
	// app.Use(recover.New()) // disable it to avoid confusion when getting runtime errors
	router.SetUpRoutes(app)

	app.Listen(":" + env.ProjectPort)
}
