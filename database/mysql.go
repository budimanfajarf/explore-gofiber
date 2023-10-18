package database

import (
	"explore-gofiber/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	MySQL *gorm.DB
)

func ConnectMySQL(env *config.IEnv) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := env.MySQLUser + ":" + env.MySQLPassword + "@tcp(" + env.MySQLHost + ":" + env.MySQLPort + ")/" + env.MySQLDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected to database")
	MySQL = db
}