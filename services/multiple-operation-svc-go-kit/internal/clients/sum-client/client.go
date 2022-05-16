package sum_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Client interface {
	Sum(ctx context.Context, request interface{}) (int, error)
}

type client struct {
	ep endpoint.Endpoint
}

func NewClient(url string) Client {
	return &client{
		ep: makeEndpoint(url),
	}
}

func (client *client) Sum(ctx context.Context, request interface{}) (int, error) {
	resp, err := client.ep(ctx, request)
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

	return httptransport.NewClient(
		"POST",
		url,
		encodeRequest,
		decodeResponse,
	).Endpoint()
}

func decodeResponse(ctx context.Context, r *http.Response) (response interface{}, err error) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	fmt.Println(r.StatusCode, buf.String())
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

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
