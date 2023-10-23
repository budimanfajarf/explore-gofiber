package article

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IRepository interface {
	FindAll(args FindAllArgs, selects []string, relations []string) ([]models.Article, error)
	FindAllAndCount(args FindAllArgs, selects []string, relations []string) ([]models.Article, int64, error)
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

func (r *repository) findAllQuery(args FindAllArgs, selects []string, relations []string) *gorm.DB {
	query := r.db.Model(&models.Article{})

	search := args.Search
	status := args.Status

	if search != "" {
		query.Where("title LIKE ?", "%"+search+"%")
	}

	if status != "" {
		query.Where("status = ?", status)
	}

	if len(selects) > 0 {
		query.Select(selects)
	}

	if len(relations) > 0 {
		for _, relation := range relations {
			if relation != "" {
				query.Preload(relation)
			}
		}
	}

	return query
}

func (r *repository) FindAll(args FindAllArgs, selects []string, relations []string) ([]models.Article, error) {
	order := utils.GetOrderValue(args.OrderBy, args.Order)
	limit := args.Limit
	offset := utils.CalculateOffset(args.Page, limit)

	query := r.findAllQuery(args, selects, relations).Order(order).Offset(offset).Limit(limit)

	var data []models.Article

	err := query.Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindAllAndCount(args FindAllArgs, selects []string, relations []string) ([]models.Article, int64, error) {
	order := utils.GetOrderValue(args.OrderBy, args.Order)
	limit := args.Limit
	offset := utils.CalculateOffset(args.Page, limit)

	dataQuery := r.findAllQuery(args, selects, relations).Order(order).Offset(offset).Limit(limit)
	countQuery := r.findAllQuery(args, selects, relations)

	var data []models.Article
	var count int64

	err := dataQuery.Find(&data).Error
	if err != nil {
		return data, count, err
	}

	err = countQuery.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (r *repository) FindOne(dest interface{}, relations []string, conds ...interface{}) *gorm.DB {
	query := r.db.Model(&models.Article{})

	if len(relations) > 0 {
		for _, relation := range relations {
			query.Preload(relation)
		}
	}

	return query.Take(dest, conds...)
}

func (r *repository) FindOneByID(dest interface{}, id uint, relations []string) *gorm.DB {
	return r.FindOne(dest, relations, "id = ?", id)
}

func (r *repository) Create(dto CreateDto) (*models.Article, error) {
	// fmt.Printf("%+v\n", dto)
	// return nil, errors.New("not implemented")

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

	// Check if tags with given IDs exist in the database
	var existingTags []models.Tag
	err = r.db.Where("id IN (?)", dto.TagIDs).Find(&existingTags).Error
	if err != nil {
		return nil, err
	}

	// Create a map to store existing tag IDs for fast lookup
	existingTagIDs := make(map[uint]bool)
	for _, tag := range existingTags {
		existingTagIDs[tag.ID] = true
	}

	var tagsToAssociate []models.Tag
	for _, id := range dto.TagIDs {
		// If the tag exists, associate it with the article
		if _, exists := existingTagIDs[id]; exists {
			tagsToAssociate = append(tagsToAssociate, models.Tag{
				Base: models.Base{
					ID: id,
				},
			})
		} else {
			// Tag with this ID doesn't exist, handle the error or situation accordingly
			return nil, fiber.NewError(400, "tag with ID "+strconv.Itoa(int(id))+" doesn't exist")
		}
	}

	// Associate existing tags with the article
	err = r.db.Model(&article).Association("Tags").Append(tagsToAssociate)
	if err != nil {
		return nil, err
	}

	return article, nil
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
