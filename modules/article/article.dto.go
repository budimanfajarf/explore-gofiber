package article

type StoreArticleDto struct {
	Title     string `validate:"required,min=5,max=150"`
	Content   string `validate:"omitempty,html"`
	Image     string `validate:"omitempty"`
	Status    string `validate:"required,oneof=UNPUBLISHED PUBLISHED"`
	CreatedBy uint
}

type GetListParams struct {
	Page   int    `json:"page" validate:"omitempty,min=1"`
	Limit  int    `json:"limit" validate:"omitempty,min=1"`
	Search string `json:"search" validate:"omitempty"`
	Status string ` validate:"omitempty,oneof=UNPUBLISHED PUBLISHED" json:"status"`
}
