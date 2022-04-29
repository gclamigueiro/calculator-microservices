package endpoint

import (
	"context"

	"github.com/gclamigueiro/sum-svc-gokit/internal/entity"
	"github.com/gclamigueiro/sum-svc-gokit/internal/service"
	"github.com/go-kit/kit/endpoint"
)

func MakeSumEndpoint(svc service.SumService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(entity.SumRequest)
		v := svc.Sum(req.A, req.B)
		return entity.SumResponse{R: v, Err: ""}, nil
	}
}
