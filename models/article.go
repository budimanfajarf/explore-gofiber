package models

// Article model
type Article struct {
	BaseModel

	Title    string `json:"title"`
	Content  string `json:"content"`
	Image    string `json:"image"`
	ImageUrl string `gorm:"-" json:"imageUrl"`
	Status   string `json:"status"`

	CreatedBy int `gorm:"column:createdBy" json:"createdBy"`
	UpdatedBy int `gorm:"column:updatedBy" json:"updatedBy"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `Article`
func (Article) TableName() string {
	return "Article"
}
