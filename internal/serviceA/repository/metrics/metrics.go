package metrics

import (
	"app/internal/metric"
)

type Metrics struct {
	RequestLatency   metric.HistogramVec
	ResponseLatency  metric.HistogramVec
	HttpRequestCount metric.CounterVec
}

func Initialize() *Metrics {
	return &Metrics{
		RequestLatency:   metric.NewHistogram(responseLatencyMetricProperties()),
		ResponseLatency:  metric.NewHistogram(requestLatencyMetricProperties()),
		HttpRequestCount: metric.NewCounter(httpRequestCountMetricProperties()),
	}
}

func responseLatencyMetricProperties() metric.Properties {
	return metric.Properties{
		Name:        "response_latency_in_seconds",
		Namespace:   "repository_a_gateway",
		Description: "Describes http response time in seconds",
		Type:        metric.HistogramVecType,
		Properties:  nil,
	}
}

func requestLatencyMetricProperties() metric.Properties {
	return metric.Properties{
		Name:        "request_latency_in_seconds",
		Namespace:   "repository_a_gateway",
		Description: "Describes http request time in seconds",
		Type:        metric.HistogramVecType,
		Properties:  nil,
	}
}

func httpRequestCountMetricProperties() metric.Properties {
	return metric.Properties{
		Name:        "http_request_count",
		Namespace:   "repository_a_gateway",
		Description: "Http requests processed",
		Type:        metric.CounterVecType,
		Properties:  nil,
	}
}
