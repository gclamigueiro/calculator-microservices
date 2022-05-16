package handler

import (
	"context"
	"encoding/json"
	"multiple-operation-svc-go-kit/internal/entity"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpHandler(sumEndpoint endpoint.Endpoint) {

	handler := httptransport.NewServer(
		sumEndpoint,
		decodeRequest,
		encodeResponse,
	)

	http.Handle("/v1/calculator/service", handler)

}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request entity.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
