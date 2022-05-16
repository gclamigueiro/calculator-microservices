package main

import (
	"net/http"
	"os"

	"github.com/go-kit/log"

	sumClient "multiple-operation-svc-go-kit/internal/clients/sum-client"
	"multiple-operation-svc-go-kit/internal/endpoint"
	"multiple-operation-svc-go-kit/internal/handler"
	"multiple-operation-svc-go-kit/internal/service"

	apiConfig "multiple-operation-svc-go-kit/cmd/config"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	apiConfig := apiConfig.GetAPIConfig()

	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "my_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "my_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "my_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	//creating clients
	sClient := sumClient.NewClient(apiConfig.UriSumService)

	svc := service.NewService(sClient)
	svc = service.NewLoggingMiddleware(logger, svc)
	svc = service.NewInstrumentingMiddleware(requestCount, requestLatency, countResult, svc)

	endp := endpoint.MakeEndpoint(svc)
	endp = endpoint.LoggingMiddleware(log.With(logger, "method", "multiple-operation"))(endp)

	handler.NewHttpHandler(endp)

	http.Handle("/metrics", promhttp.Handler())

	logger.Log("exit", http.ListenAndServe(":"+apiConfig.Port, nil))
}
