package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func CreateClinet(redisDB int) *redis.Client {

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	// Creating Redis client
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	return RedisClient
}
