package article

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
