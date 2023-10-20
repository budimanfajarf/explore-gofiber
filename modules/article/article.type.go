package article

import "time"

type ArticleListItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	ImageUrl  string    `gorm:"-" json:"imageUrl"`
	Status    string    `json:"status" validate:"required,oneof=UNPUBLISHED PUBLISHED"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
}

type StoreArticleDto struct {
	Title     string `validate:"required,min=5,max=150"`
	Content   string `validate:"omitempty,html"`
	Image     string `validate:"omitempty"`
	Status    string `validate:"required,oneof=UNPUBLISHED PUBLISHED"`
	CreatedBy uint
}

type GetListParams struct {
	Page   int    `validate:"omitempty,min=1" json:"page"`
	Limit  int    `validate:"omitempty,min=1" json:"limit"`
	Search string `validate:"omitempty" json:"search"`
	Status string `validate:"omitempty,oneof=UNPUBLISHED PUBLISHED" json:"status"`
}
