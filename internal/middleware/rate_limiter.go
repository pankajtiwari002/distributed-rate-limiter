package middleware

import (
	"distributed-rate-limiter/internal/config"
	"distributed-rate-limiter/internal/metrics"
	"log"
	"net/http"
)

type Limiter interface {
	Allow(key string, capacity int, refillRate float64) (bool, error)
}

func RateLimit(limiter Limiter, cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			clientID := r.Header.Get("X-API-KEY")
			endpoint := r.URL.Path

			key := "rate_limit:" + clientID + ":" + endpoint

			if rateLimit, exists := cfg.Limits[endpoint]; exists {
				log.Printf("Applying rate limit for client %s on endpoint %s", clientID, endpoint)
				log.Printf("Keys: %s", key)
				metrics.IncRequests()
				allowed, err := limiter.Allow(key, rateLimit.Capacity, rateLimit.RefillRate)
				if err != nil {
					if cfg.Mode == "fail-open" {
						metrics.IncFailOpen()
						next.ServeHTTP(w, r)
						return
					}
					w.WriteHeader(http.StatusTooManyRequests)
					w.Write([]byte("Too Many Requests"))
					return
				}
				if !allowed {
					metrics.IncBlocked()
					w.WriteHeader(http.StatusTooManyRequests)
					w.Write([]byte("Too Many Requests"))
					return
				}
				metrics.IncAllowed()
				next.ServeHTTP(w, r)
			} else {
				log.Printf("endpoint %s does not exist in config with clientId %s", endpoint, clientID)
				next.ServeHTTP(w, r)
			}
		})
	}
}
