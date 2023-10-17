package article

import (
	"explore-gofiber/database"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(ctx *fiber.Ctx) error {
	// // Method 1: get all columns
	// articles := []Article{}
	// database.MySQL.Find(&articles)

	// Method 2: get specific columns
	articles := []GetArticlesItemAPI{}
	database.MySQL.Model(&Article{}).Find(&articles)

	return ctx.Status(200).JSON(articles)
}
