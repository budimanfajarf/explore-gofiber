package article

import (
	"explore-gofiber/http"

	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	GetList(ctx *fiber.Ctx) error
}

type handler struct {
	service IService
}

func NewHandler(service IService) *handler {
	return &handler{
		service,
	}
}

func (h *handler) GetList(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	search := ctx.Query("search")

	data, err := h.service.GetList(page, limit, search)
	if err != nil {
		return err
	}

	meta := fiber.Map{
		"page":   page,
		"limit":  limit,
		"search": search,
	}

	return http.Success(ctx, 200, data, meta)
}
