package sum_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"multiple-operation-svc-go-kit/internal/clients/middlewares"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Client interface {
	Call(ctx context.Context, request interface{}) (int, error)
}

type client struct {
	ep endpoint.Endpoint
}

func NewClient(url string) Client {
	return &client{
		ep: makeEndpoint(url),
	}
}

func (client *client) Call(ctx context.Context, request interface{}) (int, error) {
	resp, err := client.ep(ctx, request)
	fmt.Println("SUM CLIENT", resp, err)
	if err != nil {
		return 0, err
	}
	response := resp.(*SumResponse)
	return response.R, nil
}

func makeEndpoint(uri string) endpoint.Endpoint {
	url, err := url.Parse(uri)
	if err != nil {
		fmt.Println(err)
	}

	opts := []httptransport.ClientOption{
		httptransport.SetClient(&http.Client{Timeout: 10 * time.Second}),
	}

	ep := httptransport.NewClient(
		"POST",
		url,
		encodeRequest,
		decodeResponse,
		opts...,
	).Endpoint()
	ep = middlewares.CircuitBreakerMiddleware()(ep)
	return ep
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeResponse(ctx context.Context, r *http.Response) (response interface{}, err error) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	if r.StatusCode != http.StatusOK {
		responseBody := buf.String()
		return nil, fmt.Errorf(responseBody)
	}

	resp := &SumResponse{}

	if err := json.Unmarshal(buf.Bytes(), resp); err != nil {
		return nil, err
	}

	return resp, nil
}
