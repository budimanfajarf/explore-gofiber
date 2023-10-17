package article

import (
	"explore-gofiber/common"
)

// Article model
type Article struct {
	common.BaseModel

	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `Article`
func (Article) TableName() string {
	return "Article"
}
