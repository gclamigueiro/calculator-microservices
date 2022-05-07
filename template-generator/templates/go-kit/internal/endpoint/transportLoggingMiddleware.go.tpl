package endpoint

import (
	"context"

	kitendpoint "{{.APIName}}/kit/endpoint"
	"{{.APIName}}/log"
)

type Middleware func(kitendpoint.Endpoint) kitendpoint.Endpoint

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next kitendpoint.Endpoint) kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
