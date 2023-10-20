package article

import "explore-gofiber/models"

type IService interface {
	GetList(page, limit int, search string) ([]ArticleListItem, error)
	GetDetails(id int) (*models.Article, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{
		repository,
	}
}

func (s *service) GetList(page, limit int, search string) ([]ArticleListItem, error) {
	// Test Errors
	// panic("Something went wrong")
	// return nil, fiber.NewError(fiber.StatusNotFound, "Not Found")

	return s.repository.GetList(page, limit, search)
}

func (s *service) GetDetails(id int) (*models.Article, error) {
	return s.repository.FindByID(id)
}
