package article

type IService interface {
	GetList() ([]ArticleListItem, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) *Service {
	return &Service{
		repository,
	}
}

func (s *Service) GetList() ([]ArticleListItem, error) {
	return s.repository.GetList()
}
