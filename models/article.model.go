package models

// Article model
type Article struct {
	BaseModel

	Title    string `gorm:"column:title" json:"title"`
	Content  string `gorm:"column:content" json:"content"`
	Image    string `gorm:"column:image" json:"image"`
	ImageUrl string `gorm:"-" json:"imageUrl"`
	Status   string `gorm:"column:status" json:"status"`

	Tags []Tag `gorm:"many2many:ArticleTag" json:"tags"`
}
