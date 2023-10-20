package article

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"
)

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

	data, err := s.repository.GetList(page, limit, search)
	if err != nil {
		return data, err
	}

	for i := range data {
		data[i].ImageUrl = utils.GetArticleImageURL(data[i].Image)
	}

	return data, nil
}

func (s *service) GetDetails(id int) (*models.Article, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	data.ImageUrl = utils.GetArticleImageURL(data.Image)

	return data, nil
}
