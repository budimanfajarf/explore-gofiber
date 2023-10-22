package article

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"

	"gorm.io/gorm"
)

type IRepository interface {
	GetList(params *GetListParams) ([]models.Article, error)
	FindOne(dest interface{}, relations []string, conds ...interface{}) *gorm.DB
	FindOneByID(dest interface{}, id uint, relations []string) *gorm.DB
	Create(dto CreateDto) (*models.Article, error)
	CheckIsExist(id uint) (bool, error)
	Update(id uint, dto UpdateDto) (*models.Article, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) GetList(params *GetListParams) ([]models.Article, error) {
	query := r.db.Model(&models.Article{})

	search := params.Search
	status := params.Status

	if search != "" {
		query.Where("title LIKE ?", "%"+search+"%")
	}

	if status != "" {
		query.Where("status = ?", status)
	}

	order := params.OrderBy + " " + params.Order
	limit := params.Limit
	offset := utils.CalculateOffset(params.Page, limit)

	query.Order(order).Offset(offset).Limit(limit)
	query.Select("id, title, image, status, createdAt, updatedAt")
	query.Preload("Tags")

	var data []models.Article
	err := query.Find(&data).Error

	return data, err
}

func (r *repository) FindOne(dest interface{}, relations []string, conds ...interface{}) *gorm.DB {
	query := r.db.Model(&models.Article{})

	if len(relations) > 0 {
		for _, relation := range relations {
			query = query.Preload(relation)
		}
	}

	return query.Take(dest, conds...)
}

func (r *repository) FindOneByID(dest interface{}, id uint, relations []string) *gorm.DB {
	return r.FindOne(dest, relations, "id = ?", id)
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

	return article, err
}

func (r *repository) CheckIsExist(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Article{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) Update(id uint, dto UpdateDto) (*models.Article, error) {
	article := &models.Article{
		Title:   dto.Title,
		Content: dto.Content,
		Image:   dto.Image,
		Status:  dto.Status,
		BaseModel: models.BaseModel{
			UpdatedBy: dto.UpdatedBy,
		},
	}

	err := r.db.Model(&models.Article{}).Where("id = ?", id).Updates(article).Error
	if err != nil {
		return nil, err
	}

	// to make sure the id returned
	article.ID = id

	return article, nil
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&models.Article{}, id).Error
}
