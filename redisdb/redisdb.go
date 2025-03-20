package redisdb

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// Global Redis client
var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{})
	fmt.Println("redis has been initlized!")

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
