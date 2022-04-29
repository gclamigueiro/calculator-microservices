package service

// SubtractService provides operations on strings.
type SubtractService interface {
	Subtract(A, B int) int
}

// subtractService is a concrete implementation of SubtractService
type subtractService struct{}

func NewSubtractService() SubtractService {
	return &subtractService{}
}

func (subtractService) Subtract(A, B int) int {
	return A - B
}
