package service

import (
	"context"
	subtractClient "multiple-operation-svc-go-kit/internal/clients/subtract-client"
	sumClient "multiple-operation-svc-go-kit/internal/clients/sum-client"
	"multiple-operation-svc-go-kit/internal/entity"
	"sync"
)

type Service interface {
	Call(ctx context.Context, operations []entity.Operation) ([]int, []error)
}

type service struct {
	sumClient      sumClient.Client
	subtractClient subtractClient.Client
}

func NewService(sumClient sumClient.Client, subtractClient subtractClient.Client) Service {
	return &service{
		sumClient:      sumClient,
		subtractClient: subtractClient,
	}
}

func (s *service) Call(ctx context.Context, operations []entity.Operation) ([]int, []error) {

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(operations))

	var results []int = make([]int, 0)
	var servicesErrors []error = make([]error, 0)

	resultChan := make(chan int)
	errorChan := make(chan error)

	for _, operation := range operations {
		go func(operation entity.Operation) {
			defer waitGroup.Done()
			switch operation.Operation {
			case "+":
				v, err := s.sumClient.Call(ctx, &sumClient.SumRequest{A: operation.Param1, B: operation.Param2})
				resultChan <- v
				if err != nil {
					errorChan <- err
				}
			case "-":
				v, err := s.subtractClient.Call(ctx, &subtractClient.SubtractRequest{A: operation.Param1, B: operation.Param2})
				resultChan <- v
				if err != nil {
					errorChan <- err
				}

			}
		}(operation)
	}

	go func() {
		for v := range resultChan {
			results = append(results, v)
		}
	}()

	go func() {
		for v := range errorChan {
			servicesErrors = append(servicesErrors, v)
		}
	}()

	waitGroup.Wait()
	close(resultChan)
	close(errorChan)

	return results, servicesErrors
}
