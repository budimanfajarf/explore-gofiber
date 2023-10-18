package article

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) GetList(ctx *fiber.Ctx) error {
	data, _ := h.service.GetList()

	return ctx.Status(200).JSON(data)
}
