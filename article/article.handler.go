package article

import (
	"github.com/gofiber/fiber/v2"
)

func GetArticleListHandler(ctx *fiber.Ctx) error {
	data := GetArticleList()

	return ctx.Status(200).JSON(data)
}
