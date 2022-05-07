package endpoint

import (
	"context"

	"{{.APIName}}/internal/entity"
	"{{.APIName}}/internal/service"

	"{{.APIName}}/kit/endpoint"
)

func MakeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.Request)
		v := svc.Call(req.Param1, req.Param2)
		return entity.Response{R: v, Err: ""}, nil
	}
}
