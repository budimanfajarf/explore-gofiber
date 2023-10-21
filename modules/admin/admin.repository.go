package admin

import (
	"gorm.io/gorm"
)

type IRepository interface {
	FindByEmail(email string) (*Admin, error)
	FindOne(dest interface{}, conds ...interface{}) *gorm.DB
	FindOneByEmail(dest interface{}, email string) *gorm.DB
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

func (r *repository) FindOne(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.db.Model(&Admin{}).Take(dest, conds...)
}

func (r *repository) FindOneByEmail(dest interface{}, email string) *gorm.DB {
	return r.FindOne(dest, "email = ?", email)
}
