package metrics

import "sync/atomic"

var (
	RequestsTotal   uint64
	RequestsAllowed uint64
	RequestsBlocked uint64
	RedisErrors     uint64
	FailOpenTotal   uint64
)

func IncRequests() {
	atomic.AddUint64(&RequestsTotal, 1)
}

func IncAllowed() {
	atomic.AddUint64(&RequestsAllowed, 1)
}

func IncBlocked() {
	atomic.AddUint64(&RequestsBlocked, 1)
}

func IncRedisErrors() {
	atomic.AddUint64(&RedisErrors, 1)
}

func IncFailOpen() {
	atomic.AddUint64(&FailOpenTotal, 1)
}
