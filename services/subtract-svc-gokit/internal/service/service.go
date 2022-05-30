package service

import (
	"context"
)

// SumService provides operations on strings.
type Service interface {
	Call(ctx context.Context, A, B int) (int, error)
}

// sumService is a concrete implementation of SumService
type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Call(ctx context.Context, A, B int) (int, error) {
	return A - B, nil
}
