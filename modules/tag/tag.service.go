package tag

import (
	"explore-gofiber/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type IService interface {
	FindByIDs(IDs []uint) ([]models.Tag, error)
	FindAndCheckByIDs(IDs []uint) ([]models.Tag, error)
	GetList(args FindArgs) ([]ListItem, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{
		repository,
	}
}

func (s *service) FindByIDs(IDs []uint) ([]models.Tag, error) {
	return s.repository.FindByIDs(IDs)
}

func (s *service) FindAndCheckByIDs(IDs []uint) ([]models.Tag, error) {
	tags, err := s.repository.FindByIDs(IDs)
	if err != nil {
		return tags, err
	}

	tagIDs := make(map[uint]bool)
	for _, tag := range tags {
		tagIDs[tag.ID] = true
	}

	invalidTagIDs := make([]uint, 0)
	for _, ID := range IDs {
		if _, ok := tagIDs[ID]; !ok {
			invalidTagIDs = append(invalidTagIDs, ID)
		}
	}

	if len(invalidTagIDs) > 0 {
		return tags, fiber.NewError(400, fmt.Sprintf("invalid tag IDs: %v", invalidTagIDs))
	}

	return tags, nil
}

func (s *service) GetList(args FindArgs) ([]ListItem, error) {
	var tags []ListItem
	err := s.repository.FindAll(&tags, args)
	if err != nil {
		return tags, err
	}

	return tags, nil
}
