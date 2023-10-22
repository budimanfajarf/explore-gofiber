package models

// ArticleTag model
type ArticleTag struct {
	BaseModel

	ArticleID uint `gorm:"column:articleId" json:"articleId"`
	TagID     uint `gorm:"column:tagId" json:"tagId"`
}
