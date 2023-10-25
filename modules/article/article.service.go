package article

import (
	"explore-gofiber/models"
	"explore-gofiber/modules/tag"
	"explore-gofiber/utils"
)

type IService interface {
	GetList(args FindAllArgs) ([]ListItem, int64, error)
	GetDetails(id uint) (models.Article, error)
	Create(dto CreateDto) (models.Article, error)
	Update(id uint, dto UpdateDto) (models.Article, error)
	CheckIsExist(id uint) (bool, error)
	Delete(id uint) error
}

type service struct {
	repository IRepository
	tagService tag.IService
}

func NewService(repository IRepository, tagService tag.IService) IService {
	return &service{
		repository,
		tagService,
	}
}

func (s *service) GetList(args FindAllArgs) ([]ListItem, int64, error) {
	// Test Errors
	// return nil, fiber.NewError(fiber.StatusNotFound) // caught on fiber-config.go
	// return nil, errors.New("something went wrong") // caught on fiber-config.go
	// panic("something went wrong") // caught on fiber-config.go only if enable app.Use(recover.New()) in main.go

	result := []ListItem{}

	data, count, err := s.repository.FindAllAndCount(
		args,
		[]string{"id", "title", "image", "status", "createdAt", "updatedAt"},
		[]string{"Tags"},
	)

	if err != nil {
		return result, count, err
	}

	for _, article := range data {
		item := ListItem{
			ID:        article.ID,
			Title:     article.Title,
			Image:     article.Image,
			ImageUrl:  utils.GetArticleImageURL(article.Image),
			Status:    article.Status,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
			Tags:      article.Tags,
		}
		result = append(result, item)
	}

	return result, count, nil
}

func (s *service) GetDetails(id uint) (models.Article, error) {
	data := models.Article{}
	err := s.repository.FindOneByID(&data, id, []string{"Tags"}).Error
	if err != nil {
		return data, err
	}

	data.ImageUrl = utils.GetArticleImageURL(data.Image)

	if data.Tags == nil {
		data.Tags = []models.Tag{} // make default value to empty array []
	}

	return data, nil
}

func (s *service) Create(dto CreateDto) (models.Article, error) {
	data := models.Article{}

	tags, err := s.tagService.FindAndCheckByIDs(dto.TagIDs)
	if err != nil {
		return data, err
	}

	// fmt.Printf("%+v\n", tags)

	data, err = s.repository.Create(dto, tags)
	if err != nil {
		return data, err
	}

	return s.GetDetails(data.ID)
}

func (s *service) CheckIsExist(id uint) (bool, error) {
	return s.repository.CheckIsExist(id)
}

func (s *service) Update(id uint, dto UpdateDto) (models.Article, error) {
	data := models.Article{}

	tags, err := s.tagService.FindAndCheckByIDs(dto.TagIDs)
	if err != nil {
		return data, err
	}

	data, err = s.repository.Update(id, dto, tags)
	if err != nil {
		return data, err
	}

	return s.GetDetails(data.ID)
}

func (s *service) Delete(id uint) error {
	return s.repository.Delete(id)
}
