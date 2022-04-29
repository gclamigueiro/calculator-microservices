package service

import (
	"time"

	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   SubtractService
}

func NewLoggingMiddleware(logger log.Logger, svc SubtractService) SubtractService {
	return &loggingMiddleware{logger, svc}
}

func (mw loggingMiddleware) Subtract(A, B int) (output int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", A, B,
			"output", output,
			"err", nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.next.Subtract(A, B)
	return
}
