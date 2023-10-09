package article

import (
	database "explore-gofiber/config"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(ctx *fiber.Ctx) error {
	// Method 1: get all columns
	// articles := []Article{}
	// database.DBConn.Find(&articles)

	// Method 2: get specific columns
	articles := []GetArticlesItemAPI{}
	database.DBConn.Model(&Article{}).Find(&articles)

	return ctx.Status(200).JSON(articles)
}
