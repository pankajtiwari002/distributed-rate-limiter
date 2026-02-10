# Distributed Rate Limiter (Go + Redis + Lua)

A **production-grade distributed rate limiter** built in Go using **Redis Lua scripts** for atomic enforcement across multiple API instances.  
The system supports **per-endpoint rate limits**, **fail-open / fail-closed modes**, and **observability via metrics**.

---

## Why This Project?

Rate limiting is a core backend infrastructure problem.  
This project demonstrates how real-world systems handle:

- Concurrency at scale
- Atomic state updates
- Distributed consistency
- Fault tolerance
- Config-driven behavior
- Observability

This is **not** an in-memory or single-node limiter — it is **distributed and production-ready**.

---

## High-Level Design

- **Algorithm:** Token Bucket  
- **State Store:** Redis  
- **Concurrency Control:** Redis Lua (atomic execution)  
- **Language:** Go  
- **Deployment:** Local + Docker (Windows friendly)

### Request Flow

1. Client sends request
2. Rate-limit middleware intercepts request
3. Middleware builds `(client_id + endpoint)` key
4. Go service calls Redis Lua script
5. Lua atomically refills and consumes tokens
6. Redis returns allow / block decision
7. Request proceeds or returns `429 Too Many Requests`

---

## Features

### Core
- Distributed rate limiting
- Token Bucket algorithm
- Atomic enforcement using Redis Lua
- Per-endpoint limits
- Per-user isolation
- Redis TTL cleanup

### Reliability
- Configurable **fail-open / fail-closed**
- Graceful Redis failure handling
- No local locks or mutexes

### Observability
- Prometheus-style metrics
- Allowed / blocked request counters
- Redis error tracking
- Fail-open event tracking

### Developer Friendly
- Docker-based Redis
- Clean Go module structure
- Config-driven limits

---

## Project Structure

├── cmd/
│ └── server/
│ └── main.go
├── api/
│ └── handlers/
│ └── handlers.go
├── internal/
│ ├── limiter/
│ │ └── token_bucket.go
│ ├── middleware/
│ │ └── rate_limiter.go
│ ├── redis/
│ │ ├── client.go
│ │ └── lua/
│ │ ├── token_bucket.lua
│ │ └── loader.go
│ ├── metrics/
│ │ └── metrics.go
│ └── config/
│ ├── config.go
│ └── loader.go
├── configs/
│ └── rate_limits.yaml
├── docker/
│ └── docker-compose.yml
├── go.mod
└── README.md

