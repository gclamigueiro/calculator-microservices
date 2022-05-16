package service

type SumService interface {
	Sum(A, B int) int
}

type sumService struct{}

func NewSumService() SumService {
	return &sumService{}
}

func (sumService) Sum(A, B int) int {
	return A + B
}
