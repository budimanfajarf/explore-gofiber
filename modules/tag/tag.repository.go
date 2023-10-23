package tag

import (
	"explore-gofiber/models"

	"gorm.io/gorm"
)

type IRepository interface {
	FindByIDs(IDs []uint) ([]models.Tag, error)
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
