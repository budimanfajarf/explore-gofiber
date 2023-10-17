package database

import "explore-gofiber/config"

func Connect(env *config.IEnv) {
	ConnectMySQL(env)
}
