package database

import (
	"explore-gofiber/config"
	"log"

	"github.com/gofiber/storage/redis/v3"
)

var Redis *redis.Storage

func InitRedis() {
	env := config.Env

	store := redis.New(redis.Config{
		Host:     env.RedisHost,
		Port:     env.RedisPort,
		Username: env.RedisUsername,
		Password: env.RedisPassword,
	})

	Redis = store

	// Redis.Get()

	log.Println("Redis initialized")
}
