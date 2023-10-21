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
	AdminService   admin.IService
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
	AdminService = admin.NewService(AdminRepository)
	AuthService = auth.NewService(AdminService)

	// Handlers
	ArticleHandler = article.NewHandler(ArticleService)
	AuthHandler = auth.NewHandler(AuthService)
}
