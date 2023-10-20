package article

import (
	"explore-gofiber/http"
	"explore-gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	GetList(ctx *fiber.Ctx) error
	GetDetails(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
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

func (h *handler) GetDetails(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return http.BadRequestException(ctx, "invalid article id")
	}

	// test := utils.GetArticleImageURL("test.png")
	// log.Println(test)

	article, err := h.service.GetDetails(id)
	if err != nil {
		if err.Error() == "record not found" {
			return http.NotFoundException(ctx, "article not found")
		}

		return err
	}

	return http.Success(ctx, 200, article)
}

func (h *handler) Create(ctx *fiber.Ctx) error {
	dto := new(StoreArticleDto)

	if err := utils.ParseBodyAndValidate(ctx, dto); err != nil {
		return err
	}

	return http.Success(ctx, 201, dto)
}
