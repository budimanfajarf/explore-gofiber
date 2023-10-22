package article

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"

	"gorm.io/gorm"
)

type IRepository interface {
	GetList(params *GetListParams) ([]ListItem, error)
	FindOne(dest interface{}, conds ...interface{}) *gorm.DB
	FindOneByID(dest interface{}, id uint) *gorm.DB
	FindOneByIDWithTags(dest interface{}, id uint) *gorm.DB
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

func (r *repository) GetList(params *GetListParams) ([]ListItem, error) {
	query := r.db.Model(&models.Article{})

	search := params.Search
	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	status := params.Status
	if status != "" {
		query = query.Where("status = ?", status)
	}

	orderBy := params.OrderBy
	order := params.Order
	page := params.Page
	limit := params.Limit

	query.Order(orderBy + " " + order).Offset(utils.CalculateOffset(page, limit)).Limit(limit)

	var data []ListItem
	err := query.Find(&data).Error

	return data, err
}

func (r *repository) FindOne(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.db.Model(&models.Article{}).Take(dest, conds...)
}

func (r *repository) FindOneByID(dest interface{}, id uint) *gorm.DB {
	return r.FindOne(dest, "id = ?", id)
}

func (r *repository) FindOneByIDWithTags(dest interface{}, id uint) *gorm.DB {
	return r.FindOneByID(dest, id).Preload("Tags")
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
