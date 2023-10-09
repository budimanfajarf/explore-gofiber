package article

import "time"

type GetArticlesItemAPI struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
}