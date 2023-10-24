package models

import "time"

// ArticleTag model
type ArticleTag struct {
	ArticleID uint      `gorm:"column:articleId" json:"articleId"`
	TagID     uint      `gorm:"column:tagId" json:"tagId"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
