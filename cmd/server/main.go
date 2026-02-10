package main

import (
	"log"
	"net/http"

	"distributed-rate-limiter/api/handlers"
	"distributed-rate-limiter/internal/config"
	"distributed-rate-limiter/internal/limiter"
	"distributed-rate-limiter/internal/middleware"
	redisClient "distributed-rate-limiter/internal/redis"
	"distributed-rate-limiter/internal/redis/lua"
)

func main() {
	log.Println("Starting Distributed Rate Limiter Server...")	
	// 1️⃣ Redis client
	redisClient := redisClient.New("localhost:6379")

	// 2️⃣ Load Lua script from file
	script, err := lua.LoadScript("internal/redis/lua/token_bucket.lua")
	if err != nil {
		log.Fatalf("failed to load lua script: %v", err)
	}

	// 3️⃣ Create rate limiter
	rateLimiter := limiter.NewTokenBucketLimiter(redisClient, script)

	// 4️⃣ Router
	mux := http.NewServeMux()

	// Public route
	mux.HandleFunc("/health", handlers.HealthHandler)

	rateLimitConfig, err := config.LoadConfig("config/rate_limits.yaml")
	if err != nil {
		log.Fatalf("failed to load rate limit config: %v", err)
	}
	// Rate-limited routes
	mux.Handle(
		"/api/search",
		middleware.RateLimit(rateLimiter, rateLimitConfig)(
			http.HandlerFunc(handlers.SearchHandler),
		),
	)

	mux.Handle(
		"/api/login",
		middleware.RateLimit(rateLimiter, rateLimitConfig)(
			http.HandlerFunc(handlers.LoginHandler),
		),
	)

	mux.HandleFunc("/metrics", handlers.Metrics)

	// 5️⃣ Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
