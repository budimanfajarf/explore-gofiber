package http

import "github.com/gofiber/fiber/v2"

type HttpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type HttpResponse struct {
	Data interface{} `json:"data"`
	// Meta map[string]interface{} `json:"meta"`
	// Meta  fiber.Map  `json:"meta"`
	Meta  interface{} `json:"meta"`
	Error *HttpError  `json:"error"`
}

// // func Success(ctx *fiber.Ctx, status int, data interface{}, meta ...map[string]interface{}) error {
// func Success(ctx *fiber.Ctx, status int, data interface{}, meta ...fiber.Map) error {
// 	var metaData fiber.Map

// 	if len(meta) > 0 {
// 		metaData = meta[0]
// 	}

// 	return ctx.Status(status).JSON(HttpResponse{
// 		Data: data,
// 		Meta: metaData,
// 	})
// }

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
