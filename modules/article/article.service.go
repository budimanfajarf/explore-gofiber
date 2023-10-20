package article

// import "github.com/gofiber/fiber/v2"

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
	// Test Errors
	// panic("Something went wrong")
	// return nil, fiber.NewError(fiber.StatusNotFound, "Not Found")

	return s.repository.GetList()
}
