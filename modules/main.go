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
	mySqlDB := database.MySQL

	// Repositories
	ArticleRepository = article.NewRepository(mySqlDB)

	// Services
	ArticleService = article.NewService(ArticleRepository)

	// Handlers
	ArticleHandler = article.NewHandler(ArticleService)
}
