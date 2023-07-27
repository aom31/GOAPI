package database

import (
	"log"
	"project-for-portfolioDEV/GOAPI/config"

	"github.com/go-redis/redis/v8"
)

func ConnectionRedisDb(config *config.Config) *redis.Client {

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})
	log.Println("Success Connected to Redis with ", config.RedisUrl)

	return redisClient
}
