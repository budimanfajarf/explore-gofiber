package admin

import (
	"gorm.io/gorm"
)

type IRepository interface {
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

func (r *repository) FindOne(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.db.Model(&Admin{}).Take(dest, conds...)
}

func (r *repository) FindOneByEmail(dest interface{}, email string) *gorm.DB {
	return r.FindOne(dest, "email = ?", email)
}
