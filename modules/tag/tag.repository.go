package tag

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"
	"fmt"

	"gorm.io/gorm"
)

type IRepository interface {
	FindByIDs(IDs []uint) ([]models.Tag, error)
	FindAll(args FindAllArgs) ([]ListItem, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) FindByIDs(IDs []uint) ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Where("id IN (?)", IDs).Find(&tags).Error
	if err != nil {
		return tags, err
	}

	return tags, nil
}

func (r *repository) FindAll(args FindAllArgs) ([]ListItem, error) {
	order := utils.GetOrderValue(args.OrderBy, args.Order)

	var tags []ListItem
	err := r.db.Model(&models.Tag{}).Order(order).Find(&tags).Error
	fmt.Printf("tags: %+v", tags)

	if err != nil {
		return tags, err
	}

	return tags, nil
}
