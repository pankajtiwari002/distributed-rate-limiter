package handlers

import (
	"fmt"
	"net/http"

	"distributed-rate-limiter/internal/metrics"
)

func Metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"requests_total %d\n"+
			"requests_allowed %d\n"+
			"requests_blocked %d\n"+
			"redis_errors %d\n"+
			"fail_open_total %d\n",
		metrics.RequestsTotal,
		metrics.RequestsAllowed,
		metrics.RequestsBlocked,
		metrics.RedisErrors,
		metrics.FailOpenTotal,
	)
}
