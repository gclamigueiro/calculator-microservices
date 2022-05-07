package main

import (
	"net/http"
	"os"

	"{{.APIName}}/log"

	"{{.APIName}}/internal/endpoint"
	"{{.APIName}}/internal/handler"
	"{{.APIName}}/internal/service"

	kitprometheus "{{.APIName}}/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	svc := service.NewSumService()
	svc = service.NewLoggingMiddleware(logger, svc)
	svc = service.NewInstrumentingMiddleware(requestCount, requestLatency, countResult, svc)

	endp := endpoint.MakeEndpoint(svc)
	endp = endpoint.LoggingMiddleware(log.With(logger, "method", "sum"))(endp)

	handler.NewHttpHandler(endp)

	http.Handle("/metrics", promhttp.Handler())

	logger.Log("exit", http.ListenAndServe(":8080", nil))
}