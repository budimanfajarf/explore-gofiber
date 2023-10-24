package tag

import (
	"explore-gofiber/utils/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service IService
}

func NewHandler(router fiber.Router, service IService) {
	handler := &handler{
		service,
	}

	router.Get("/", handler.getList)
}

func (h *handler) getList(ctx *fiber.Ctx) error {
	params := FindAllArgs{
		OrderBy: ctx.Query("orderBy", "id"),
		Order:   ctx.Query("order", "desc"),
	}

	tags, err := h.service.GetList(params)
	if err != nil {
		return err
	}

	return http.ResponseWithMeta(ctx, 200, tags, params)
}
