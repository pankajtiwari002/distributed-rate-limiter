package limiter

type RateLimiter interface {
	Allow(key string, capacity int, refillRate float64) (bool, error)
}
