package middlewares

import (
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "Clients"
	st.Timeout = 30
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.3
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

type Middleware func(kitendpoint.Endpoint) kitendpoint.Endpoint

func CircuitBreakerMiddleware() Middleware {
	return func(next kitendpoint.Endpoint) kitendpoint.Endpoint {
		next = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))(next)
		next = circuitbreaker.Gobreaker(cb)(next)
		return next
	}
}
