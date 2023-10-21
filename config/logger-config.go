package config

import "github.com/gofiber/fiber/v2/middleware/logger"

func LoggerConfig() logger.Config {
	return logger.Config{
		// Format:     "${time} | ${ip} | ${method} | ${path} | ${status} | ${latency}\n",
		TimeFormat: "2006/01/02 15:04:05",
	}
}
