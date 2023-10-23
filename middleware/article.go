package middleware

import (
	"explore-gofiber/modules"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CheckIfArticleExist(ctx *fiber.Ctx) error {
	articleRepository := modules.ArticleRepository

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(400, "invalid article id")
	}

	isArticleExist, err := articleRepository.CheckIsExist(uint(id))
	if err != nil {
		return err
	}

	if !isArticleExist {
		return fiber.NewError(404, fmt.Sprintf("article with id %d not exist", id))
	}

	return ctx.Next()
}
