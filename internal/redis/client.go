package redisClient

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	rdb *redis.Client
}

func New(addr string) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &Client{rdb: rdb}
}

func (c *Client) Eval(
	ctx context.Context,
	script string,
	keys []string,
	args ...interface{},
) (interface{}, error) {
	return c.rdb.Eval(ctx, script, keys, args...).Result()
}
