package endpoint

import (
	"context"

	"github.com/gclamigueiro/subtract-svc-gokit/internal/entity"
	"github.com/gclamigueiro/subtract-svc-gokit/internal/service"
	"github.com/go-kit/kit/endpoint"
)

func MakeSumEndpoint(svc service.SubtractService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.SubtractRequest)
		v := svc.Subtract(req.A, req.B)
		return entity.SubtractResponse{R: v, Err: ""}, nil
	}
}
