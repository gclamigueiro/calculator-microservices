package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           Service
}

func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram,
	countResult metrics.Histogram, svc Service) Service {
	return &instrumentingMiddleware{requestCount, requestLatency, countResult, svc}
}

func (mw instrumentingMiddleware) Call(ctx context.Context, Param1, Param2 int) (output int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "sum", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output = mw.next.Call(ctx, Param1, Param2)
	return
}
