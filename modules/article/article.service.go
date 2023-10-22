package article

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"
	// "github.com/gofiber/fiber/v2"
)

type IService interface {
	GetList(params *GetListParams) ([]ListItem, error)
	GetDetails(id uint) (*models.Article, error)
	Create(dto CreateDto) (*models.Article, error)
	Update(id uint, dto UpdateDto) (*models.Article, error)
	Delete(id uint) error
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{
		repository,
	}
}

func (s *service) GetList(params *GetListParams) ([]ListItem, error) {
	// Test Errors
	// return nil, fiber.NewError(fiber.StatusNotFound) // caught on fiber-config.go
	// return nil, errors.New("something went wrong") // caught on fiber-config.go
	// panic("something went wrong") // caught on fiber-config.go only if enable app.Use(recover.New()) in main.go

	data, err := s.repository.GetList(params)
	if err != nil {
		return data, err
	}

	for i := range data {
		data[i].ImageUrl = utils.GetArticleImageURL(data[i].Image)
	}

	return data, nil
}

func (s *service) GetDetails(id uint) (*models.Article, error) {
	data := &models.Article{}
	err := s.repository.FindOneByIDWithTags(data, id).Error
	if err != nil {
		return data, err
	}

	data.ImageUrl = utils.GetArticleImageURL(data.Image)

	return data, nil
}

func (s *service) Create(dto CreateDto) (*models.Article, error) {
	data, err := s.repository.Create(dto)
	if err != nil {
		return nil, err
	}

	return s.GetDetails(data.ID)
}

func (s *service) Update(id uint, dto UpdateDto) (*models.Article, error) {
	data, err := s.repository.Update(id, dto)
	if err != nil {
		return nil, err
	}

	return s.GetDetails(data.ID)
}

func (s *service) Delete(id uint) error {
	return s.repository.Delete(id)
}
