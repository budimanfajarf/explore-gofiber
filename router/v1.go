package router

import (
	"explore-gofiber/middleware"
	"explore-gofiber/modules"

	"github.com/gofiber/fiber/v2"
)

func setUpRoutesV1(app *fiber.App) {
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	// articles
	v1Article := v1.Group("/articles")
	v1Article.Get("/", modules.ArticleHandler.GetList)
	v1Article.Get("/:id", modules.ArticleHandler.GetDetails)
	v1Article.Post("/", modules.ArticleHandler.Create)
	v1Article.Put("/:id", middleware.IsArticleExistMiddleware, modules.ArticleHandler.Update)
	v1Article.Delete("/:id", middleware.IsArticleExistMiddleware, modules.ArticleHandler.Delete)

	// auth
	v1Auth := v1.Group("/auth")
	v1Auth.Post("/login", modules.AuthHandler.Login)
}