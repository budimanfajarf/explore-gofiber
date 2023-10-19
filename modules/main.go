package modules

import (
	"explore-gofiber/database"
	"explore-gofiber/modules/article"
)

var (
	// Repositories
	articleRepository article.IRepository

	// Services
	articleService article.IService

	// Handlers
	ArticleHandler article.IHandler
)

func Init() {
	mySqlDB := database.MySQL

	// Repositories
	articleRepository = article.NewRepository(mySqlDB)

	// Services
	articleService = article.NewService(articleRepository)

	// Handlers
	ArticleHandler = article.NewHandler(articleService)
}
