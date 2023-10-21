package modules

import (
	"explore-gofiber/database"
	"explore-gofiber/modules/admin"
	"explore-gofiber/modules/article"
)

var (
	// Repositories
	ArticleRepository article.IRepository
	AdminRepository   admin.IRepository

	// Services
	ArticleService article.IService
	AdminService   admin.IService

	// Handlers
	ArticleHandler article.IHandler
)

func Init() {
	mySqlDB := database.MySQL

	// Repositories
	ArticleRepository = article.NewRepository(mySqlDB)
	AdminRepository = admin.NewRepository(mySqlDB)

	// Services
	ArticleService = article.NewService(ArticleRepository)
	AdminService = admin.NewService(AdminRepository)

	// Handlers
	ArticleHandler = article.NewHandler(ArticleService)
}
