package article

type StoreArticleDto struct {
	Title     string `validate:"required,min=5,max=150"`
	Content   string `validate:"required,html"`
	Image     string `validate:""`
	Status    string `json:"status" validate:"required,oneof=UNPUBLISHED PUBLISHED"`
	CreatedBy uint
}
