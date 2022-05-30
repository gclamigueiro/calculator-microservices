package service

import (
	"context"
)

// SumService provides operations on strings.
type Service interface {
	Call(ctx context.Context, Param1, Param2 int) (int, error)
}

// sumService is a concrete implementation of SumService
type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Call(ctx context.Context, Param1, Param2 int) (int, error) {
	// Do some work
	return 0, nil
}
