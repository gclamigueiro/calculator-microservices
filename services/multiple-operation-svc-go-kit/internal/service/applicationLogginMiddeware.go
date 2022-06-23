package service

import (
	"context"
	"time"

	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

func NewLoggingMiddleware(logger log.Logger, svc Service) Service {
	return &loggingMiddleware{logger, svc}
}

func (mw loggingMiddleware) Call(ctx context.Context, expression string) (output int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "multiple operations",
			"input", expression,
			"output", output,
			"err", nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Call(ctx, expression)
	return
}
