package article

import (
	"explore-gofiber/utils/scopes"

	"gorm.io/gorm"
)

func searchScope(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			return db.Where("title LIKE ?", "%"+search+"%")
		}
		return db
	}
}

func statusScope(status string) func(db *gorm.DB) *gorm.DB {
	return scopes.StatusScope(status)
}

func orderScope(orderBy string, order string) func(db *gorm.DB) *gorm.DB {
	return scopes.OrderScope(orderBy, order)
}

func paginationScope(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return scopes.PaginationScope(page, limit)
}

func selectScope(selects []string) func(db *gorm.DB) *gorm.DB {
	return scopes.SelectScope(selects)
}

func relationsScope(relations []string) func(db *gorm.DB) *gorm.DB {
	return scopes.RelationsScope(relations)
}
