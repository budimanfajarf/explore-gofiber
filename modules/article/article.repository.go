package article

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"

	"gorm.io/gorm"
)

type IRepository interface {
	GetList(page, limit int, search string) ([]ArticleListItem, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) GetList(page, limit int, search string) ([]ArticleListItem, error) {
	var data []ArticleListItem

	query := r.db.Model(&models.Article{})

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	query.Limit(limit).Offset(utils.CalculateOffset(page, limit))

	err := query.Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
