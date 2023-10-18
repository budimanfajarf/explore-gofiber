package article

import (
	"explore-gofiber/models"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	GetList() ([]ArticleListItem, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{
		db,
	}
}

func (r *articleRepository) GetList() ([]ArticleListItem, error) {
	var data []ArticleListItem
	err := r.db.Model(&models.Article{}).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
