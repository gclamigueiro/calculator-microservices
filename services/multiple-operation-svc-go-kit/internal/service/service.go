package service

import (
	"context"
	subtractClient "multiple-operation-svc-go-kit/internal/clients/subtract-client"
	sumClient "multiple-operation-svc-go-kit/internal/clients/sum-client"
	"strconv"
)

var supportedOperations = map[string]bool{
	"+": true,
	"-": true,
}

type Service interface {
	Call(ctx context.Context, expression string) (int, error)
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

func (s *service) Call(ctx context.Context, expression string) (int, error) {

	values := make([]int, 0)
	ops := make([]string, 0)

	number := ""

	for _, r := range expression {
		v := string(r)
		_, isOperation := supportedOperations[v]
		if isOperation {
			ops = append(ops, v)
			if number != "" {
				n, _ := strconv.Atoi(number)
				values = append(values, n)
				number = ""
			}
		} else {
			number = number + v
		}
	}

	if number != "" {
		n, _ := strconv.Atoi(number)
		values = append(values, n)
	}

	for len(values) > 1 {
		v1, v2 := 0, 0
		v1, v2, values = values[0], values[1], values[2:]

		op := ops[0]
		ops = ops[1:]

		switch op {
		case "+":
			r, _ := s.sumClient.Call(ctx, &sumClient.SumRequest{A: v1, B: v2})
			values = append([]int{r}, values...)
		case "-":
			r, _ := s.subtractClient.Call(ctx, &subtractClient.SubtractRequest{A: v1, B: v2})
			values = append([]int{r}, values...)
		}

	}

	/*var waitGroup sync.WaitGroup
	var results []int = make([]int, 0)
	var servicesErrors []error = make([]error, 0)

	type responseStruct struct {
		result int
		err    error
	}

	resultChan := make(chan *responseStruct)

	for _, operation := range operations {
		waitGroup.Add(1)
		go func(operation entity.Operation) {

			switch operation.Operation {
			case "+":
				v, err := s.sumClient.Call(ctx, &sumClient.SumRequest{A: operation.Param1, B: operation.Param2})
				resultChan <- &responseStruct{
					result: v,
					err:    err,
				}

			case "-":
				v, err := s.subtractClient.Call(ctx, &subtractClient.SubtractRequest{A: operation.Param1, B: operation.Param2})
				resultChan <- &responseStruct{
					result: v,
					err:    err,
				}
			}
			waitGroup.Done()
		}(operation)
	}

	go func() {
		waitGroup.Wait()
		close(resultChan)
	}()

	for v := range resultChan {
		results = append(results, v.result)
		if v.err != nil {
			servicesErrors = append(servicesErrors, v.err)
		}
	}
	*/
	return values[0], nil
}
