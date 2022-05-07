package service

// SumService provides operations on strings.
type Service interface {
	Call(Param1, Param2 int) int
}

// sumService is a concrete implementation of SumService
type service struct{}

func NewSumService() Service {
	return &service{}
}

func (service) Call(Param1, Param2 int) int {
	// Do some work
	return 0
}
