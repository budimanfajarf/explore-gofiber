package config

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GormConfig = &gorm.Config{
	Logger: logger.Default.LogMode(logger.Info),
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
		NoLowerCase:   true,
	},
}
