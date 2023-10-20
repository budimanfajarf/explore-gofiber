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
	data, err := h.service.GetList()
	if err != nil {
		return err
	}

	return http.Success(ctx, 200, data)
}
