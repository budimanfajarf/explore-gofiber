package article

type IArticleService interface {
	GetList() ([]ArticleListItem, error)
}

type articleService struct {
	articleRepository IArticleRepository
}

func NewArticleService(articleRepository IArticleRepository) *articleService {
	return &articleService{
		articleRepository,
	}
}

func (s *articleService) GetList() ([]ArticleListItem, error) {
	return s.articleRepository.GetList()
}
