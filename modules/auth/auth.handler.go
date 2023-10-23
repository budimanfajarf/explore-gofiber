package auth

import (
	"explore-gofiber/http"
	"explore-gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	Login(ctx *fiber.Ctx) error
}

type handler struct {
	service IService
}

func NewHandler(service IService) *handler {
	return &handler{
		service,
	}
}

func (h *handler) Login(ctx *fiber.Ctx) error {
	dto := new(LoginDto)
	if err := utils.ParseBodyAndValidate(ctx, dto); err != nil {
		return err
	}

	data, err := h.service.Login(*dto)
	if err != nil {
		return err
	}

	return http.Response(ctx, 200, data)
}
