package article

import "github.com/gofiber/fiber/v2"

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
	data, _ := h.service.GetList()

	return ctx.Status(200).JSON(data)
}
