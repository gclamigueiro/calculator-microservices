package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gclamigueiro/sum-svc-gokit/internal/entity"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpHandler(sumEndpoint endpoint.Endpoint) {

	sumHandler := httptransport.NewServer(
		sumEndpoint,
		decodeRequest,
		encodeResponse,
	)

	http.Handle("/", sumHandler)

}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request entity.SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
