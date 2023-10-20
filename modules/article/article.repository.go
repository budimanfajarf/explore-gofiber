package article

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"

	"gorm.io/gorm"
)

type IRepository interface {
	GetList(params *GetListParams) ([]ListItem, error)
	FindByID(id uint) (*models.Article, error)
	Create(dto CreateDto) (*models.Article, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) GetList(params *GetListParams) ([]ListItem, error) {
	page := params.Page
	limit := params.Limit
	search := params.Search
	status := params.Status

	query := r.db.Model(&models.Article{})

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Limit(limit).Offset(utils.CalculateOffset(page, limit))

	var data []ListItem
	err := query.Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindByID(id uint) (*models.Article, error) {
	var data models.Article

	err := r.db.Model(&models.Article{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) Create(dto CreateDto) (*models.Article, error) {
	article := &models.Article{
		Title:   dto.Title,
		Content: dto.Content,
		Image:   dto.Image,
		Status:  dto.Status,
		BaseModel: models.BaseModel{
			CreatedBy: dto.CreatedBy,
			UpdatedBy: dto.CreatedBy,
		},
	}

	err := r.db.Create(article).Error
	if err != nil {
		return nil, err
	}

	return article, nil
}
