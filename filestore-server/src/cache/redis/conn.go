package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
)

var redisClient *redis.Client
var ctx = context.Background()

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        "0.0.0.0:6380",
		DB:          0,
		PoolSize:    30,
		IdleTimeout: time.Minute * 2,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("redis connect err ", err)
	}
}

// GetClient get redis client
func GetClient() *redis.Client {
	return redisClient
}
