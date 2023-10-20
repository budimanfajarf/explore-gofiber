package article

import "time"

type ArticleListItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	ImageUrl  string    `gorm:"-" json:"imageUrl"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
}
