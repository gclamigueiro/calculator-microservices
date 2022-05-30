package endpoint

import (
	"context"

	"github.com/gclamigueiro/subtract-svc-gokit/internal/entity"
	"github.com/gclamigueiro/subtract-svc-gokit/internal/service"
	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Request)
		v, err := svc.Call(ctx, req.A, req.B)
		if err != nil {
			return entity.Response{R: 0, Err: err.Error()}, nil
		}
		return entity.Response{R: v, Err: ""}, nil
	}
}
