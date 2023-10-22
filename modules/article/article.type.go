package article

import (
	"explore-gofiber/models"
	"time"
)

type ListItem struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Image     string       `json:"image"`
	ImageURL  string       `json:"imageUrl"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	Tags      []models.Tag `json:"tags"`
}

type CreateDto struct {
	Title     string `validate:"required,min=5,max=150"`
	Content   string `validate:"omitempty,html"`
	Image     string `validate:"omitempty"`
	Status    string `validate:"required,oneof=UNPUBLISHED PUBLISHED"`
	CreatedBy uint
}

type GetListParams struct {
	Page    int    `validate:"omitempty,min=1" default:"1" json:"page"`
	Limit   int    `validate:"omitempty,min=1" json:"limit"`
	OrderBy string `validate:"omitempty,oneof=id title createdAt updatedAt image status" json:"orderBy"`
	Order   string `validate:"omitempty,oneof=asc desc" json:"order"`
	Search  string `validate:"omitempty" json:"search"`
	Status  string `validate:"omitempty,oneof=UNPUBLISHED PUBLISHED" json:"status"`
}

type UpdateDto struct {
	Title     string `validate:"required,min=5,max=150"`
	Content   string `validate:"omitempty,html"`
	Image     string `validate:"omitempty"`
	Status    string `validate:"required,oneof=UNPUBLISHED PUBLISHED"`
	UpdatedBy uint
}
