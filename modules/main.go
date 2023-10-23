package modules

import (
	"explore-gofiber/database"
	"explore-gofiber/modules/admin"
	"explore-gofiber/modules/article"
	"explore-gofiber/modules/auth"
	"explore-gofiber/modules/tag"
)

var (
	// Repositories
	ArticleRepository article.IRepository
	AdminRepository   admin.IRepository
	TagRepository     tag.IRepository

	// Services
	ArticleService article.IService
	AuthService    auth.IService
	TagService     tag.IService

	// Handlers
	ArticleHandler article.IHandler
	AuthHandler    auth.IHandler
)

func Init() {
	mySqlDB := database.GormMySqlDBConn

	// Repositories
	ArticleRepository = article.NewRepository(mySqlDB)
	AdminRepository = admin.NewRepository(mySqlDB)
	TagRepository = tag.NewRepository(mySqlDB)

	// Services
	TagService = tag.NewService(TagRepository)
	ArticleService = article.NewService(ArticleRepository, TagService)
	AuthService = auth.NewService(AdminRepository)

	// Handlers
	ArticleHandler = article.NewHandler(ArticleService)
	AuthHandler = auth.NewHandler(AuthService)
}
