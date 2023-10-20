package article

import (
	// "errors"
	"explore-gofiber/models"
	"explore-gofiber/utils"
	// "github.com/gofiber/fiber/v2"
)

type IService interface {
	GetList(page, limit int, search string) ([]ArticleListItem, error)
	GetDetails(id uint) (*models.Article, error)
	Create(dto StoreArticleDto) (*models.Article, error)
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
	// return nil, fiber.NewError(fiber.StatusNotFound) // caught on fiber-config.go
	// return nil, errors.New("something went wrong") // caught on fiber-config.go
	// panic("something went wrong") // caught on fiber-config.go only if enable app.Use(recover.New()) in main.go

	data, err := s.repository.GetList(page, limit, search)
	if err != nil {
		return data, err
	}

	for i := range data {
		data[i].ImageUrl = utils.GetArticleImageURL(data[i].Image)
	}

	return data, nil
}

func (s *service) GetDetails(id uint) (*models.Article, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	data.ImageUrl = utils.GetArticleImageURL(data.Image)

	return data, nil
}

func (s *service) Create(dto StoreArticleDto) (*models.Article, error) {
	data, err := s.repository.Create(dto)
	if err != nil {
		return nil, err
	}

	return data, nil
}
