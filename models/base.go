package models

import "time"

type Base struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

type BaseModel struct {
	Base

	CreatedBy int `gorm:"column:createdBy" json:"createdBy"`
	UpdatedBy int `gorm:"column:updatedBy" json:"updatedBy"`
}
