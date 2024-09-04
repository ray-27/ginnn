package db

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8" // Import the Redis library
)

var ctx = context.Background()

func Redis_init() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use "localhost:6379" if running locally without password
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	// Testing connectivity
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong, "Connected to Redis")

	return rdb
}


