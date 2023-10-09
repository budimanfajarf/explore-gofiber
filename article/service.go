package article

import (
	database "explore-gofiber/config"

	"github.com/gofiber/fiber/v2"
)

func GetArticles(ctx *fiber.Ctx) error {
	articles := []Article{}

	err := database.DBConn.Find(&articles).Error
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(articles)
}
