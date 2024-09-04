package services

import (
	"context"
	"gin/services/auth"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client
var ctx = context.Background()

func SetClient(rdb *redis.Client) {
	client = rdb
	println(auth.AuthFun())
}
