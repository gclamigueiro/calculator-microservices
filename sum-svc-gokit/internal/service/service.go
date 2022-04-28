package service

// SumService provides operations on strings.
type SumService interface {
	Sum(A, B int) int
}

// sumService is a concrete implementation of SumService
type sumService struct{}

func NewSumService() SumService {
	return &sumService{}
}

func (sumService) Sum(A, B int) int {
	return A + B
}
