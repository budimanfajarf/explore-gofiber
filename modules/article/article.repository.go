package article

import (
	"explore-gofiber/models"

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
	offset := (page - 1) * limit

	query := r.db.Model(&models.Article{}).Limit(limit).Offset(offset)

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	err := query.Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
