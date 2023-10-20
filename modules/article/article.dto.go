package article

type StoreArticleDto struct {
	Title  string `validate:"required,min=5,max=150"`
	Status string `validate:"required,oneof=UNPUBLISHED PUBLISHED"`
}
