package admin

import (
	"gorm.io/gorm"
)

type IRepository interface {
	FindByEmail(email string) (*Admin, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) FindByEmail(email string) (*Admin, error) {
	var data Admin
	err := r.db.Model(&Admin{}).Where("email = ?", email).First(&data).Error
	return &data, err
}
