package article

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"
	"time"
)

type ListItem struct {
	ID        uint         `gorm:"column:id;primarykey" json:"id"`
	Title     string       `gorm:"column:title" json:"title"`
	Image     string       `gorm:"column:image" json:"image"`
	ImageUrl  string       `gorm:"-" json:"imageUrl"`
	Status    string       `gorm:"column:status" json:"status"`
	CreatedAt time.Time    `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time    `gorm:"column:updatedAt" json:"updatedAt"`
	Tags      []models.Tag `json:"tags"`
}

type CreateDto struct {
	Title     string `validate:"required,min=5,max=150" json:"title"`
	Content   string `validate:"omitempty,html" json:"content"`
	Image     string `validate:"omitempty" json:"image"`
	Status    string `validate:"required,oneof=UNPUBLISHED PUBLISHED" json:"status"`
	TagIDs    []uint `validate:"omitempty" json:"tagIDs"`
	CreatedBy uint
}

type FindAllArgs struct {
	Page    int    `validate:"omitempty,min=1" default:"1" json:"page"`
	Limit   int    `validate:"omitempty,min=1" json:"limit"`
	OrderBy string `validate:"omitempty,oneof=id title createdAt updatedAt image status" json:"orderBy"`
	Order   string `validate:"omitempty,oneof=asc desc" json:"order"`
	Search  string `validate:"omitempty" json:"search"`
	Status  string `validate:"omitempty,oneof=UNPUBLISHED PUBLISHED" json:"status"`
}

type UpdateDto struct {
	Title     string `validate:"required,min=5,max=150" json:"title"`
	Content   string `validate:"omitempty,html" json:"content"`
	Image     string `validate:"omitempty" json:"image"`
	Status    string `validate:"required,oneof=UNPUBLISHED PUBLISHED" json:"status"`
	TagIDs    []uint `validate:"omitempty" json:"tagIDs"`
	UpdatedBy uint
}

// extended from utils.PaginationMeta
type GetListMeta struct {
	Count     int64                  `json:"count"`
	Page      int                    `json:"page"`
	Limit     int                    `json:"limit"`
	TotalPage int                    `json:"totalPage"`
	PrevPage  *int                   `json:"prevPage"`
	NextPage  *int                   `json:"nextPage"`
	From      int                    `json:"from"`
	To        int                    `json:"to"`
	Links     []utils.PaginationLink `json:"links"`
	OrderBy   string                 `json:"orderBy"`
	Order     string                 `json:"order"`
	Search    string                 `json:"search"`
	Status    string                 `json:"status"`
}
