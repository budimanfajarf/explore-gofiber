package middleware

import (
	"explore-gofiber/http"
	"explore-gofiber/modules"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func IsArticleExistMiddleware(ctx *fiber.Ctx) error {
	articleRepository := modules.ArticleRepository

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return http.BadRequestException(ctx, "invalid article id")
	}

	isArticleExist, err := articleRepository.CheckIsExist(uint(id))
	if err != nil {
		return err
	}

	if !isArticleExist {
		return http.NotFoundException(ctx, fmt.Sprintf("article with id %d not exist", id))
	}

	return ctx.Next()
}
