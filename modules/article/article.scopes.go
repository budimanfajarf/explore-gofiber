package article

import (
	"explore-gofiber/utils/scopes"

	"gorm.io/gorm"
)

func SearchScope(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			return db.Where("title LIKE ?", "%"+search+"%")
		}
		return db
	}
}

func StatusScope(status string) func(db *gorm.DB) *gorm.DB {
	return scopes.StatusScope(status)
}

func OrderScope(orderBy string, order string) func(db *gorm.DB) *gorm.DB {
	return scopes.OrderScope(orderBy, order)
}

func PaginationScope(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return scopes.PaginationScope(page, limit)
}

func SelectScope(selects []string) func(db *gorm.DB) *gorm.DB {
	return scopes.SelectScope(selects)
}

func RelationsScope(relations []string) func(db *gorm.DB) *gorm.DB {
	return scopes.RelationsScope(relations)
}
