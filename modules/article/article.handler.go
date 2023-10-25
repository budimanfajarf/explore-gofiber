package article

import (
	"explore-gofiber/utils"
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
	router.Get("/:id", handler.getDetails)
	router.Post("/", handler.create)
	router.Put("/:id", handler.checkIsExistMiddleware, handler.update)
	router.Delete("/:id", handler.checkIsExistMiddleware, handler.delete)
}

func (h *handler) getList(ctx *fiber.Ctx) error {
	params := FindArgs{
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

	return http.ResponseWithMeta(ctx, 200, data, GetListMeta{
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

func (h *handler) getDetails(ctx *fiber.Ctx) error {
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

	return http.Response(ctx, 200, article)
}

func (h *handler) create(ctx *fiber.Ctx) error {
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

	return http.Response(ctx, 201, data)
}

func (h *handler) update(ctx *fiber.Ctx) error {
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

	return http.Response(ctx, 200, data)
}

func (h *handler) delete(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	err := h.service.Delete(uint(id))
	if err != nil {
		return err
	}

	return http.Response(ctx, 200, "article deleted successfully")
}
