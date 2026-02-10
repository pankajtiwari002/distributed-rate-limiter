package limiter

import (
	"context"
	"time"
)

type RedisClient interface {
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error)
}

type TokenBucketLimiter struct {
	redis RedisClient
	script string
}

func NewTokenBucketLimiter(redis RedisClient, script string) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		redis: redis,
		script: script,
	}
}

func (l *TokenBucketLimiter) Allow(
	key string,
	capacity int,
	refillRate float64,
) (bool, error) {

	now := time.Now().Unix()

	result, err := l.redis.Eval(
		context.Background(),
		l.script,
		[]string{key},
		capacity,
		refillRate,
		now,
	)

	if err != nil {
		return false, err
	}

	return result.(int64) == 1, nil
}
