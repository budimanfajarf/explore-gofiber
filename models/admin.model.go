package models

import "time"

// Admin model
type Admin struct {
	ID        uint      `gorm:"column:id;primarykey" json:"id"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	CreatedBy uint      `gorm:"column:createdBy" json:"createdBy"`
	UpdatedBy uint      `gorm:"column:updatedBy" json:"updatedBy"`
}

func (Admin) TableName() string {
	return "Admin"
}
