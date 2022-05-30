package main

import (
	"net/http"
	"os"

	"github.com/go-kit/log"

	"{{.APINamespace}}{{.APIName}}/internal/endpoint"
	"{{.APINamespace}}{{.APIName}}/internal/handler"
	"{{.APINamespace}}{{.APIName}}/internal/service"
	apiConfig "{{.APINamespace}}{{.APIName}}/cmd/config"

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
	}, []string{}) // no fields here

	svc := service.NewService()
	svc = service.NewLoggingMiddleware(logger, svc)
	svc = service.NewInstrumentingMiddleware(requestCount, requestLatency, countResult, svc)

	endp := endpoint.MakeEndpoint(svc)
	endp = endpoint.LoggingMiddleware(log.With(logger, "method", "sum"))(endp)

	handler.NewHttpHandler(endp)

	http.Handle("/metrics", promhttp.Handler())

	logger.Log("exit", http.ListenAndServe(":"+apiConfig.Port, nil))
}
