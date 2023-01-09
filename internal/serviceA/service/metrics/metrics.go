package metrics

import (
	"app/internal/metric"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	RequestLatency   prometheus.Collector
	ResponseLatency  prometheus.Collector
	HttpRequestCount prometheus.Collector
}

func Initialize() *Metrics {
	return &Metrics{
		RequestLatency:   metric.New(newResponseLatencyMetrics()),
		ResponseLatency:  metric.New(newRequestLatencyMetrics()),
		HttpRequestCount: metric.New(newHttpRequestCountMetrics()),
	}
}

func newResponseLatencyMetrics() metric.Properties {
	return metric.Properties{
		Name:        "response_latency_in_seconds",
		Namespace:   "service_a_gateway",
		Description: "Describes http response time in seconds",
		Type:        metric.HistogramVec,
		Properties:  nil,
	}
}

func newRequestLatencyMetrics() metric.Properties {
	return metric.Properties{
		Name:        "request_latency_in_seconds",
		Namespace:   "service_a_gateway",
		Description: "Describes http request time in seconds",
		Type:        metric.HistogramVec,
		Properties:  nil,
	}
}

func newHttpRequestCountMetrics() metric.Properties {
	return metric.Properties{
		Name:        "http_request_count",
		Namespace:   "service_a_gateway",
		Description: "Http requests processed",
		Type:        metric.CounterVec,
		Properties:  nil,
	}
}
