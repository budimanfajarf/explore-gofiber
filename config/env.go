package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ProjectPort string
}

func LoadEnv() (*EnvConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	config := &EnvConfig{
		ProjectPort: os.Getenv("PROJECT_PORT"),
	}

	return config, nil
}
