package models

import "time"

// Article model
type Article struct {
	ID        uint      `gorm:"column:id;primarykey" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	Content   string    `gorm:"column:content" json:"content"`
	Image     string    `gorm:"column:image" json:"image"`
	ImageUrl  string    `gorm:"-" json:"imageUrl"`
	Status    string    `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	CreatedBy uint      `gorm:"column:createdBy" json:"createdBy"`
	UpdatedBy uint      `gorm:"column:updatedBy" json:"updatedBy"`

	Tags []Tag `gorm:"many2many:ArticleTag" json:"tags"`
}
