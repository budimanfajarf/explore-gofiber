package auth

import (
	"explore-gofiber/http"
	"explore-gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service IService
}

func NewHandler(router fiber.Router, service IService) {
	handler := &handler{
		service,
	}

	router.Post("/login", handler.Login)
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
