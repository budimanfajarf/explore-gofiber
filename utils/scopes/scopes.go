// common gorm scopes
package scopes

import (
	"fmt"

	"gorm.io/gorm"
)

func StatusScope(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			return db.Where("status = ?", status)
		}
		return db
	}
}

func OrderScope(orderBy string, order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if orderBy != "" && order != "" {
			return db.Order(fmt.Sprintf("%s %s", orderBy, order))
		}
		return db
	}
}

func PaginationScope(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * limit).Limit(limit)
	}
}

func SelectScope(selects []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(selects) > 0 {
			return db.Select(selects)
		}
		return db
	}
}

func RelationsScope(relations []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(relations) > 0 {
			for _, relation := range relations {
				if relation != "" {
					db.Preload(relation)
				}
			}
		}
		return db
	}
}
