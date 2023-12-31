package tag

import (
	"explore-gofiber/utils/scopes"

	"gorm.io/gorm"
)

func orderScope(orderBy string, order string) func(db *gorm.DB) *gorm.DB {
	return scopes.OrderScope(orderBy, order)
}
