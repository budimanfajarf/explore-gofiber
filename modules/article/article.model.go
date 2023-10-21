package article

import "explore-gofiber/types"

// Article model
type Article struct {
	types.BaseModel

	Title    string `json:"title"`
	Content  string `json:"content"`
	Image    string `json:"image"`
	ImageUrl string `gorm:"-" json:"imageUrl"`
	Status   string `json:"status"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `Article`
func (Article) TableName() string {
	return "Article"
}
