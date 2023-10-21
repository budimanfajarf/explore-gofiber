package modules

import (
	"explore-gofiber/database"
	"explore-gofiber/modules/admin"
	"explore-gofiber/modules/article"
	"explore-gofiber/modules/auth"
)

var (
	// Repositories
	ArticleRepository article.IRepository
	AdminRepository   admin.IRepository

	// Services
	ArticleService article.IService
	AuthService    auth.IService

	// Handlers
	ArticleHandler article.IHandler
	AuthHandler    auth.IHandler
)

func Init() {
	mySqlDB := database.MySQL

	// Repositories
	ArticleRepository = article.NewRepository(mySqlDB)
	AdminRepository = admin.NewRepository(mySqlDB)

	// Services
	ArticleService = article.NewService(ArticleRepository)
	AuthService = auth.NewService(AdminRepository)

	// Handlers
	ArticleHandler = article.NewHandler(ArticleService)
	AuthHandler = auth.NewHandler(AuthService)
}
