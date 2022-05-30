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
		v, err := svc.Call(ctx, req.Operations)

		errorsString := []string{}

		for _, e := range err {
			errorsString = append(errorsString, e.Error())
		}

		return entity.Response{R: v, Err: errorsString}, nil
	}
}
