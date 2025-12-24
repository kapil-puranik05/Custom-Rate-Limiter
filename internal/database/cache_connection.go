package database

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx context.Context

func ConnectCache() {
	opt, err := redis.ParseURL("")
	if err != nil {
		log.Println("Unable to connect to the cache")
		return
	}
	Ctx = context.Background()
	Client = redis.NewClient(opt)
}
