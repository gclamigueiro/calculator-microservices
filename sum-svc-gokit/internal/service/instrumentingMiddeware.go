package service

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           SumService
}

func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram,
	countResult metrics.Histogram, svc SumService) SumService {
	return &instrumentingMiddleware{requestCount, requestLatency, countResult, svc}
}

func (mw instrumentingMiddleware) Sum(A, B int) (output int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "sum", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output = mw.next.Sum(A, B)
	return
}
