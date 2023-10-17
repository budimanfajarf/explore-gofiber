package models

import "time"

type Base struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

type BaseModel struct {
	Base

	CreatedBy uint `gorm:"column:createdBy" json:"createdBy"`
	UpdatedBy uint `gorm:"column:updatedBy" json:"updatedBy"`
}

type BaseWithSoftDelete struct {
	Base

	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
}

type BaseModelWithSoftDelete struct {
	BaseModel

	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
}
