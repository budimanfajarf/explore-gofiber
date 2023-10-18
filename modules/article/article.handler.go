package article

import "github.com/gofiber/fiber/v2"

type articleHandler struct {
	articleService IArticleService
}

func NewArticleHandler(articleService IArticleService) *articleHandler {
	return &articleHandler{
		articleService,
	}
}

func (h *articleHandler) GetList(ctx *fiber.Ctx) error {
	data, _ := h.articleService.GetList()

	return ctx.Status(200).JSON(data)
}
