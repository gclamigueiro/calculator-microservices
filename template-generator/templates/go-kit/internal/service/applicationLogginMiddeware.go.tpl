package service

import (
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

func (mw loggingMiddleware) Call(Param1, Param2 int) (output int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "call",
			"input", Param1, Param2,
			"output", output,
			"err", nil,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.next.Call(Param1, Param2)
	return
}
