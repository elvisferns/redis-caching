package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheAPI interface {
	Get(ctx context.Context, key string, data interface{}) (result Result, err error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (result Result, err error)
	Delete(ctx context.Context, key string) (result Result, err error)
	Close()
}

func (c *cache) Get(ctx context.Context, key string) (result Result, err error) {
	val, rerr := c.rClient.Get(ctx, key).Result()
	var res Result = Result{Val: val}
	// var err error;
	switch {
	case rerr == redis.Nil:
		err = MyError{
			When: time.Now(),
			What: fmt.Sprintf("Key: %s not available", key),
		}

	case rerr != nil:
		err = MyError{
			When: time.Now(),
			What: fmt.Sprintf("Error getting key: %s :: %s", key, rerr),
		}
	}

	return res, err
}

func (c *cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (result Result, err error) {
	val, rerr := c.rClient.Set(ctx, key, value, expiration).Result()
	var res Result = Result{Val: val}

	if rerr != nil {
		err = MyError{
			When: time.Now(),
			What: fmt.Sprintf("Error setting key: %s : %s", key, rerr),
		}
	}

	return res, err
}

func (c *cache) Delete(ctx context.Context, key string) (result Result, err error) {
	val, rerr := c.rClient.Del(ctx, key).Result()
	var res Result = Result{Val: val}

	if rerr != nil {
		err = MyError{
			When: time.Now(),
			What: fmt.Sprintf("Error deleting key: %s", key),
		}
	}

	if val == 0 {
		err = MyError{
			When: time.Now(),
			What: "Key does not exist",
		}
	}

	return res, err
}

func (c *cache) Close() (err error) {
	return c.rClient.Close()
}
