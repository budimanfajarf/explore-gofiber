package models

// Article model
type Article struct {
	BaseModel

	Title    string `json:"title"`
	Content  string `json:"content"`
	Image    string `json:"image"`
	ImageUrl string `gorm:"-" json:"imageUrl"`
	Status   string `json:"status"`
}

func (Article) TableName() string {
	return "Article"
}
