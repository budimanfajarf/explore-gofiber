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
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
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
	params := FindAllArgs{
		Page:    ctx.QueryInt("page", 1),
		Limit:   ctx.QueryInt("limit", 10),
		OrderBy: ctx.Query("orderBy", "id"),
		Order:   ctx.Query("order", "desc"),
		Search:  ctx.Query("search"),
		Status:  ctx.Query("status"),
	}

	if err := utils.Validate(params); err != nil {
		return err
	}

	data, count, err := h.service.GetList(params)
	if err != nil {
		return err
	}

	paginationMeta := utils.GeneratePaginationMeta(count, params.Page, params.Limit)

	return http.SuccessWithMeta(ctx, 200, data, GetListMeta{
		Count:     paginationMeta.Count,
		Page:      paginationMeta.Page,
		Limit:     paginationMeta.Limit,
		TotalPage: paginationMeta.TotalPage,
		PrevPage:  paginationMeta.PrevPage,
		NextPage:  paginationMeta.NextPage,
		From:      paginationMeta.From,
		To:        paginationMeta.To,
		Links:     paginationMeta.Links,
		OrderBy:   params.OrderBy,
		Order:     params.Order,
		Search:    params.Search,
		Status:    params.Status,
	})
}

func (h *handler) GetDetails(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(400, "invalid article id")
	}

	article, err := h.service.GetDetails(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return fiber.NewError(404, "article not found")
		}

		return err
	}

	return http.Success(ctx, 200, article)
}

func (h *handler) Create(ctx *fiber.Ctx) error {
	dto := new(CreateDto)

	if err := utils.ParseBodyAndValidate(ctx, dto); err != nil {
		return fiber.NewError(400, err.Error())
	}

	authenticatedUserId := utils.GetUser(ctx)
	dto.CreatedBy = authenticatedUserId

	data, err := h.service.Create(*dto)
	if err != nil {
		return err
	}

	return http.Success(ctx, 201, data)
}

func (h *handler) Update(ctx *fiber.Ctx) error {
	dto := new(UpdateDto)

	if err := utils.ParseBodyAndValidate(ctx, dto); err != nil {
		return fiber.NewError(400, err.Error())
	}

	id, _ := ctx.ParamsInt("id") // no need to check error, already checked on CheckIfArticleExist middleware
	dto.UpdatedBy = utils.GetUser(ctx)

	data, err := h.service.Update(uint(id), *dto)
	if err != nil {
		return err
	}

	return http.Success(ctx, 200, data)
}

func (h *handler) Delete(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	err := h.service.Delete(uint(id))
	if err != nil {
		return err
	}

	return http.Success(ctx, 200, "article deleted successfully")
}
