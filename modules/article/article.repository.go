package article

import (
	"explore-gofiber/models"

	"gorm.io/gorm"
)

type IRepository interface {
	GetList() ([]ArticleListItem, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) GetList() ([]ArticleListItem, error) {
	var data []ArticleListItem
	err := r.db.Model(&models.Article{}).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
