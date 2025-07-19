package config

import (
	"github.com/redis/go-redis/v9"
)

type redisConfig struct {
	address  string
	password string
}

func getRedis(config redisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.address,
		Password: config.password,
	})
	return client
}
