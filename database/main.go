package database

import "explore-gofiber/config"

func Connect(env *config.EnvConfig) {
	ConnectMySQL(env)
}
