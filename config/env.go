package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ProjectPort string
}

func LoadEnv() *EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(2)
	}

	config := &EnvConfig{
		ProjectPort: os.Getenv("PROJECT_PORT"),
	}

	return config
}
