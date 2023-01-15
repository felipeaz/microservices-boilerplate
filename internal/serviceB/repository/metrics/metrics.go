package metrics

import (
	"app/internal/metric"
)

const (
	NameProperty        = "repository_b_latency_in_seconds"
	NamespaceProperty   = "repository_b_gateway"
	DescriptionProperty = "Describes http response time in seconds"

	queryTypePropertyKey = "queryType"
)

type Metrics struct {
	Latency metric.HistogramVec
}

func Initialize() *Metrics {
	return &Metrics{
		Latency: metric.NewHistogram(latencyMetricProperties()),
	}
}

func latencyMetricProperties() metric.Properties {
	return metric.Properties{
		Name:        NameProperty,
		Namespace:   NamespaceProperty,
		Description: DescriptionProperty,
		Type:        metric.HistogramVecType,
		Properties:  []string{queryTypePropertyKey},
	}
}
