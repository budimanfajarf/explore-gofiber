package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDb() {
	envConfig := LoadEnv()

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := envConfig.MySQLUser + ":" + envConfig.MySQLPassword + "@tcp(" + envConfig.MySQLHost + ":" + envConfig.MySQLPort + ")/" + envConfig.MySQLDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database")
	DBConn = db
}
