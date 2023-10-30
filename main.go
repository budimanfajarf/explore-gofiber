package main

import (
	"explore-gofiber/config"
	"explore-gofiber/database"
	"explore-gofiber/modules"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	env := config.LoadEnv()
	database.Connect()

	app := fiber.New(config.FiberConfig)
	app.Use(logger.New(config.LoggerConfig))
	// app.Use(recover.New()) // disable it to avoid confusion when getting runtime errors

	app.Use(swagger.New(swagger.Config{
		Next:     nil,
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Explore Go Fiber API documentation",
		// CacheAge: 3600, // Default to 1 hour
		CacheAge: 0,
	}))

	modules.Init(app)

	app.Listen(":" + env.ProjectPort)
}
