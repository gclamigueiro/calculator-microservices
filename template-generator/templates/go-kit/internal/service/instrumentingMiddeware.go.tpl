package service

import (
	"time"

	"{{.APIName}}/kit/metrics"
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

func (mw instrumentingMiddleware) Call(Param1, Param2 int) (output int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "sum", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output = mw.next.Call(Param1, Param2)
	return
}
