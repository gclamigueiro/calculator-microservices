package service

import (
	"context"
	"multiple-operation-svc-go-kit/internal/entity"
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

func (mw loggingMiddleware) Call(ctx context.Context, operations []entity.Operation) (output []int, err []error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "multiple operations",
			"input", operations,
			"output", output,
			"err", nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Call(ctx, operations)
	return
}
