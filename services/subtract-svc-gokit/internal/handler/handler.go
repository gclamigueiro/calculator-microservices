package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/gclamigueiro/subtract-svc-gokit/internal/entity"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpHandler(sumEndpoint endpoint.Endpoint) {

	handler := httptransport.NewServer(
		sumEndpoint,
		decodeRequest,
		encodeResponse,
	)

	http.Handle("/", handler)

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
