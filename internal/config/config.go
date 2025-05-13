package config

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAddr string
	Port 			int
}

func Load() (*Config, error) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 5000
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	return &Config{
		RedisAddr: redisAddr,
		Port: port,
	}, nil
}