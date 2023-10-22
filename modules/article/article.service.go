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

	var result []ListItem

	data, err := s.repository.GetList(params)
	if err != nil {
		return result, err
	}

	for _, article := range data {
		item := ListItem{
			ID:        article.ID,
			Title:     article.Title,
			Image:     article.Image,
			ImageURL:  utils.GetArticleImageURL(article.Image),
			Status:    article.Status,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
			Tags:      article.Tags,
		}
		result = append(result, item)
	}

	return result, nil
}

func (s *service) GetDetails(id uint) (*models.Article, error) {
	data := &models.Article{}
	// err := s.repository.FindOne(data, nil, "id = ?", id).Error
	// err := s.repository.FindOne(data, []string{"Tags"}, "id = ?", id).Error
	// err := s.repository.FindOneByID(data, id, nil).Error
	err := s.repository.FindOneByID(data, id, []string{"Tags"}).Error
	if err != nil {
		return data, err
	}

	data.ImageUrl = utils.GetArticleImageURL(data.Image)

	if data.Tags == nil {
		data.Tags = []models.Tag{}
	}

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
