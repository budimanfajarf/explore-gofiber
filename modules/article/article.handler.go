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
	search := ctx.Query("search")
	page := ctx.QueryInt("page", 1)

	data, err := h.service.GetList()
	if err != nil {
		return err
	}

	meta := fiber.Map{
		"search": search,
		"page":   page,
	}

	return http.Success(ctx, 200, data, meta)
}
