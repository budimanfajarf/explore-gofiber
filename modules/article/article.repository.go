package article

import (
	"explore-gofiber/types"
	"explore-gofiber/utils"

	"gorm.io/gorm"
)

type IRepository interface {
	GetList(params *GetListParams) ([]ListItem, error)
	FindByID(id uint) (*Article, error)
	Create(dto CreateDto) (*Article, error)
	CheckIsExist(id uint) (bool, error)
	Update(id uint, dto UpdateDto) (*Article, error)
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
	query := r.db.Model(&Article{})

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

func (r *repository) FindByID(id uint) (*Article, error) {
	var data Article

	err := r.db.Model(&Article{}).Where("id = ?", id).First(&data).Error

	return &data, err
}

func (r *repository) Create(dto CreateDto) (*Article, error) {
	article := &Article{
		Title:   dto.Title,
		Content: dto.Content,
		Image:   dto.Image,
		Status:  dto.Status,
		BaseModel: types.BaseModel{
			CreatedBy: dto.CreatedBy,
			UpdatedBy: dto.CreatedBy,
		},
	}

	err := r.db.Create(article).Error

	return article, err
}

func (r *repository) CheckIsExist(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&Article{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) Update(id uint, dto UpdateDto) (*Article, error) {
	article := &Article{
		Title:   dto.Title,
		Content: dto.Content,
		Image:   dto.Image,
		Status:  dto.Status,
		BaseModel: types.BaseModel{
			UpdatedBy: dto.UpdatedBy,
		},
	}

	err := r.db.Model(&Article{}).Where("id = ?", id).Updates(article).Error
	if err != nil {
		return nil, err
	}

	// to make sure the id returned
	article.ID = id

	return article, nil
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Article{}, id).Error
}
