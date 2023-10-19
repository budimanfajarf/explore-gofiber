package modules

import (
	"explore-gofiber/database"
	"explore-gofiber/modules/article"
)

var (
	// Repositories
	ArticleRepository article.IRepository

	// Services
	ArticleService article.IService

	// Handlers
	ArticleHandler article.IHandler
)

func Init() {
	db := database.MySQL

	// Repositories
	ArticleRepository = article.NewRepository(db)

	// Services
	ArticleService = article.NewService(ArticleRepository)

	// Handlers
	ArticleHandler = article.NewHandler(ArticleService)
}
