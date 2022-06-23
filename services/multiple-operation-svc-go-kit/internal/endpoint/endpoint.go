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
		v, err := svc.Call(ctx, req.Expression)

		if err != nil {
			return nil, err
		}

		return entity.Response{R: v}, nil
	}
}
