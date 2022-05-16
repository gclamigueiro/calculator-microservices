package service

import (
	"context"
	"fmt"
	sumClient "multiple-operation-svc-go-kit/internal/clients/sum-client"
)

type Service interface {
	Call(ctx context.Context, Param1, Param2 int) int
}

type service struct {
	sumClient sumClient.Client
}

func NewService(sumClient sumClient.Client) Service {
	return &service{
		sumClient: sumClient,
	}
}

func (s *service) Call(ctx context.Context, Param1, Param2 int) int {

	v, err := s.sumClient.Sum(ctx, &sumClient.SumRequest{A: Param1, B: Param2})

	if err != nil {
		fmt.Println(err)
	}

	return v
}
