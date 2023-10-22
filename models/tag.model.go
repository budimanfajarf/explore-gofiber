package models

// Tag model
type Tag struct {
	Base

	Name string `gorm:"column:name" json:"name"`
}
