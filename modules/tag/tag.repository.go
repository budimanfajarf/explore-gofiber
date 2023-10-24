package tag

import (
	"explore-gofiber/models"
	"explore-gofiber/utils"

	"gorm.io/gorm"
)

type IRepository interface {
	FindByIDs(IDs []uint) ([]models.Tag, error)
	FindAll(result interface{}, args FindAllArgs) error
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

func (r *repository) FindAll(result interface{}, args FindAllArgs) error {
	order := utils.GetOrderValue(args.OrderBy, args.Order)
	// note:
	// don't add "&" in result, e.g "&result", it will throw stack error
	// because already add "&" on the service layer (see tag.service.go->GetList)
	err := r.db.Model(models.Tag{}).Order(order).Find(result).Error
	// fmt.Printf("result: %+v", result)
	return err
}
