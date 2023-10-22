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
	// page := ctx.QueryInt("page", 1)
	// limit := ctx.QueryInt("limit", 10)
	// search := ctx.Query("search")
	// status := ctx.Query("status")

	// if status != "" && status != "UNPUBLISHED" && status != "PUBLISHED" {
	// 	return http.BadRequestException(ctx, "invalid status, status should be UNPUBLISHED or PUBLISHED")
	// }

	params := &GetListParams{
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

	data, err := h.service.GetList(params)
	if err != nil {
		return err
	}

	// Create a custom response slice with the desired fields
	var response []ListItem
	for _, article := range data {
		item := ListItem{
			ID:        article.ID,
			Title:     article.Title,
			Image:     article.Image,
			ImageURL:  utils.GetArticleImageURL(article.Image),
			Status:    article.Status,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
			Tags:      article.Tags,
		}
		response = append(response, item)
	}

	return http.SuccessWithMeta(ctx, 200, response, params)
}

func (h *handler) GetDetails(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return http.BadRequestException(ctx, "invalid article id")
	}

	// test := utils.GetArticleImageURL("test.png")
	// log.Println(test)

	article, err := h.service.GetDetails(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			return http.NotFoundException(ctx, "article not found")
		}

		return err
	}

	return http.Success(ctx, 200, article)
}

func (h *handler) Create(ctx *fiber.Ctx) error {
	dto := new(CreateDto)

	if err := utils.ParseBodyAndValidate(ctx, dto); err != nil {
		return http.InvalidPayloadException(ctx, err.Error())
	}

	authenticatedUserId := *utils.GetUser(ctx)
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
		return http.InvalidPayloadException(ctx, err.Error())
	}

	id, _ := ctx.ParamsInt("id") // no need to check error, already checked on CheckIfArticleExist middleware
	dto.UpdatedBy = *utils.GetUser(ctx)

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
