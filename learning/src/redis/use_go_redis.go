package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        "0.0.0.0:6379",
		DB:          0,
		PoolSize:    10,
		IdleTimeout: 100,
	})

	pong, err := redisClient.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func goRedisGetSet() {
	defer redisClient.Close()

	err := redisClient.Set(ctx, "redis", "go-redis", 0).Err()
	if err != nil {
		fmt.Println("set err", err)
		return
	}

	val1, err := redisClient.Get(ctx, "redis").Result()
	if err != nil {
		fmt.Println("get err", err)
		return
	}
	fmt.Println("get redis", val1)

	err = redisClient.HSet(ctx, "myhash", map[string]interface{}{"key1": "val1", "key2": "val2"}).Err()
	if err != nil {
		fmt.Println("hset err", err)
		return
	}

	val2, err := redisClient.HGetAll(ctx, "myhash").Result()
	if err != nil {
		fmt.Println("hgetall err", err)
		return
	}
	fmt.Println(val2)

	exist, err := redisClient.HExists(ctx, "myhash", "key1").Result()
	fmt.Println("val1 is exist?", exist)

	err = redisClient.Expire(ctx, "myhash", 10*time.Second).Err()
}
