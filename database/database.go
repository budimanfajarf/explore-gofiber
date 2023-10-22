package database

import (
	"explore-gofiber/config"
	"explore-gofiber/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	GormMySqlDBConn *gorm.DB
)

func Connect() {
	env := config.Env
	connectMySQL(env)
}

func connectMySQL(env *config.IEnv) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := env.MySQLUser + ":" + env.MySQLPassword + "@tcp(" + env.MySQLHost + ":" + env.MySQLPort + ")/" + env.MySQLDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/
	db, err := gorm.Open(mysql.Open(dsn), config.GormConfig)
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	err = db.SetupJoinTable(&models.Article{}, "Tags", &models.ArticleTag{})
	if err != nil {
		log.Fatal("Failed to setup join table. \n", err)
	}

	GormMySqlDBConn = db
	log.Println("Connected to database")
}
