package article

import (
	"explore-gofiber/models"
)

type ListItem struct {
	models.Base

	Title    string `json:"title"`
	Image    string `json:"image"`
	ImageUrl string `gorm:"-" json:"imageUrl"`
	Status   string `json:"status"`
}

type CreateDto struct {
	Title     string `validate:"required,min=5,max=150"`
	Content   string `validate:"omitempty,html"`
	Image     string `validate:"omitempty"`
	Status    string `validate:"required,oneof=UNPUBLISHED PUBLISHED"`
	CreatedBy uint
}

type GetListParams struct {
	Page    int    `validate:"omitempty,min=1" json:"page"`
	Limit   int    `validate:"omitempty,min=1" json:"limit"`
	Search  string `validate:"omitempty" json:"search"`
	Status  string `validate:"omitempty,oneof=UNPUBLISHED PUBLISHED" json:"status"`
	OrderBy string `validate:"omitempty,oneof=id title createdAt updatedAt image status" json:"orderBy"`
	Order   string `validate:"omitempty,oneof=asc desc" json:"order"`
}
