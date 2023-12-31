package article

import (
	"explore-gofiber/models"

	"gorm.io/gorm"
)

type IRepository interface {
	FindAndCount(args FindArgs, selects []string, relations []string) ([]models.Article, int64, error)
	FindOne(dest interface{}, relations []string, conds ...interface{}) *gorm.DB
	FindOneByID(dest interface{}, id uint, relations []string) *gorm.DB
	Create(dto CreateDto, tags []models.Tag) (models.Article, error)
	CheckIsExist(id uint) (bool, error)
	Update(id uint, dto UpdateDto, tags []models.Tag) (models.Article, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &repository{
		db,
	}
}

func (r *repository) FindAndCount(args FindArgs, selects []string, relations []string) ([]models.Article, int64, error) {
	var data []models.Article
	var count int64

	dataQuery := r.db.Model(&models.Article{}).Scopes(
		statusScope(args.Status),
		searchScope(args.Search),
		orderScope(args.OrderBy, args.Order),
		paginationScope(args.Page, args.Limit),
		selectScope(selects),
		relationsScope(relations),
	)

	err := dataQuery.Find(&data).Error
	if err != nil {
		return data, count, err
	}

	countQuery := r.db.Model(&models.Article{}).Scopes(
		statusScope(args.Status),
		searchScope(args.Search),
	)

	err = countQuery.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (r *repository) FindOne(dest interface{}, relations []string, conds ...interface{}) *gorm.DB {
	return r.db.Model(&models.Article{}).Scopes(relationsScope(relations)).Take(dest, conds...)
}

func (r *repository) FindOneByID(dest interface{}, id uint, relations []string) *gorm.DB {
	return r.FindOne(dest, relations, "id = ?", id)
}

func (r *repository) Create(dto CreateDto, tags []models.Tag) (models.Article, error) {
	// fmt.Printf("%+v\n", dto)
	// return nil, errors.New("not implemented")

	article := models.Article{
		Title:     dto.Title,
		Content:   dto.Content,
		Image:     dto.Image,
		Status:    dto.Status,
		CreatedBy: dto.CreatedBy,
		UpdatedBy: dto.CreatedBy,
	}

	// err := r.db.Create(&article).Error
	// if err != nil {
	// 	return article, err
	// }

	// err = r.db.Model(&article).Association("Tags").Append(tags)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&article).Error; err != nil {
			return err
		}

		if err := tx.Model(&article).Association("Tags").Append(tags); err != nil {
			return err
		}

		// return errors.New("something went wrong on transaction") // test error on transaction, when error appeared, db should rollback and not create article and associate tags
		return nil
	})

	if err != nil {
		return article, err
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

func (r *repository) Update(id uint, dto UpdateDto, tags []models.Tag) (models.Article, error) {
	article := models.Article{
		ID:        id,
		Title:     dto.Title,
		Content:   dto.Content,
		Image:     dto.Image,
		Status:    dto.Status,
		UpdatedBy: dto.UpdatedBy,
	}

	// err := r.db.Model(&models.Article{}).Where("id = ?", id).Updates(&article).Error
	// if err != nil {
	// 	return article, err
	// }

	// fmt.Printf("%+v\n", article)

	// err = r.db.Model(&article).Association("Tags").Replace(tags)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Article{}).Where("id = ?", id).Updates(&article).Error; err != nil {
			return err
		}

		// // It sometimes duplicating relation table items, the "ArticleTag"
		// if err := tx.Model(&article).Association("Tags").Replace(tags); err != nil {
		// 	return err
		// }

		if err := tx.Model(&article).Association("Tags").Clear(); err != nil {
			return err
		}

		if err := tx.Model(&article).Association("Tags").Append(tags); err != nil {
			return err
		}

		// return errors.New("something went wrong on transaction") // test error on transaction, when error appeared, db should rollback and not update the article and not replace the associate tags
		return nil
	})

	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&models.Article{}, id).Error
}
