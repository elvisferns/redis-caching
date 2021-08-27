package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type MyError struct {
	// test tag 1
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

type cache struct {
	rClient *redis.Client
}

type Result struct {
	Val interface{}
}

func NewClient(ctx context.Context) (redisCache cache, err error) {
	redisCache.rClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err = redisCache.rClient.Ping(ctx).Result()
	return redisCache, err
}
