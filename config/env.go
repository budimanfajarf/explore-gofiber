package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IEnv struct {
	ProjectPort   string
	MySQLHost     string
	MySQLPort     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string
	StorageUrl    string
}

var Env *IEnv

func LoadEnv() *IEnv {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// MySQLPort, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	// if err != nil {
	// 	panic(err)
	// }

	Env = &IEnv{
		ProjectPort: os.Getenv("PROJECT_PORT"),
		MySQLHost:   os.Getenv("MYSQL_HOST"),
		// MySQLPort:     MySQLPort,
		MySQLPort:     os.Getenv("MYSQL_PORT"),
		MySQLUser:     os.Getenv("MYSQL_USER"),
		MySQLPassword: os.Getenv("MYSQL_PASSWORD"),
		MySQLDatabase: os.Getenv("MYSQL_DB_NAME"),
		StorageUrl:    os.Getenv("STORAGE_URL"),
	}

	return Env
}
