package endpoint

import (
	"context"
	"{{.APINamespace}}{{.APIName}}/internal/entity"
	"{{.APINamespace}}{{.APIName}}/internal/service"
	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Request)
		v, err := svc.Call(ctx, req.Param1, req.Param2)
		if(err != nil) {
			return entity.Response{R: 0, Err: err.Error()}, nil
		}
		return entity.Response{R: v, Err: ""}, nil
	}
}
