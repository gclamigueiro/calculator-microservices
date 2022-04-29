package service

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           SubtractService
}

func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram,
	countResult metrics.Histogram, svc SubtractService) SubtractService {
	return &instrumentingMiddleware{requestCount, requestLatency, countResult, svc}
}

func (mw instrumentingMiddleware) Subtract(A, B int) (output int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "sum", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output = mw.next.Subtract(A, B)
	return
}
