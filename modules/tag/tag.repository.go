package tag

import (
	"explore-gofiber/models"

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
	// note:
	// don't add "&" in result, e.g "&result", it will throw stack error
	// because already add "&" on the service layer (see tag.service.go->GetList)
	err := r.db.Model(models.Tag{}).Scopes(OrderScope(args.OrderBy, args.Order)).Find(result).Error
	// fmt.Printf("result: %+v", result)
	return err
}
