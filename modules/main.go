package modules

import (
	"explore-gofiber/database"
	"explore-gofiber/modules/admin"
	"explore-gofiber/modules/article"
	"explore-gofiber/modules/auth"
	"explore-gofiber/modules/tag"
	"explore-gofiber/utils/http"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	mySqlDB := database.GormMySqlDBConn

	// Admin
	adminRepository := admin.NewRepository(mySqlDB)

	// Auth
	authService := auth.NewService(adminRepository)

	// Tag
	tagRepository := tag.NewRepository(mySqlDB)
	tagService := tag.NewService(tagRepository)

	// Article
	articleRepository := article.NewRepository(mySqlDB)
	articleService := article.NewService(articleRepository, tagService)

	// Routes & Handlers
	app.Get("/", func(c *fiber.Ctx) error {
		return http.Response(c, 200, "Hello World")
	})

	v1 := app.Group("/v1")

	auth.NewHandler(v1.Group("/auth"), authService)
	article.NewHandler(v1.Group("/articles"), articleService)

	app.Use(func(c *fiber.Ctx) error {
		return fiber.NewError(404)
	})
}
