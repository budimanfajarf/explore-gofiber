package article

import (
	database "explore-gofiber/config"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(ctx *fiber.Ctx) error {
	articles := []Article{}

	database.DBConn.Find(&articles)

	return ctx.Status(200).JSON(articles)
}
