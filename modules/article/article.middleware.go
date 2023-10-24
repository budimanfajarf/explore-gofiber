package article

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) checkIsExistMiddleware(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(400, "invalid article id")
	}

	isArticleExist, err := h.service.CheckIsExist(uint(id))
	if err != nil {
		return err
	}

	if !isArticleExist {
		return fiber.NewError(404, fmt.Sprintf("article with id %d not exist", id))
	}

	return ctx.Next()
}
