package endpoint

import (
	"context"
	"multiple-operation-svc-go-kit/internal/entity"
	"multiple-operation-svc-go-kit/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Request)
		v := svc.Call(ctx, req.Param1, req.Param2)
		return entity.Response{R: v, Err: ""}, nil
	}
}
