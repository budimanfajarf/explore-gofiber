package http

import "github.com/gofiber/fiber/v2"

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HttpResponse struct {
	Data  interface{} `json:"data"`
	Meta  interface{} `json:"meta"`
	Error *HttpError  `json:"error"`
}

func Success(ctx *fiber.Ctx, status int, data interface{}) error {
	return ctx.Status(status).JSON(HttpResponse{
		Data: data,
	})
}

func SuccessWithMeta(ctx *fiber.Ctx, status int, data interface{}, meta interface{}) error {
	return ctx.Status(status).JSON(HttpResponse{
		Data: data,
		Meta: meta,
	})
}
