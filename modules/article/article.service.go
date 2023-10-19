package article

type IService interface {
	GetList() ([]ArticleListItem, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{
		repository,
	}
}

func (s *service) GetList() ([]ArticleListItem, error) {
	return s.repository.GetList()
}
