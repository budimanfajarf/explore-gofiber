package admin

type IService interface {
	FindByEmail(email string) (*Admin, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{
		repository,
	}
}

func (s *service) FindByEmail(email string) (*Admin, error) {
	data, err := s.repository.FindByEmail(email)
	return data, err
}
