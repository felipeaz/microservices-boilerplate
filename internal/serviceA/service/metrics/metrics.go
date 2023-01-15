package metrics

import (
	"app/internal/metric"
)

const (
	NameProperty        = "error_count"
	NamespaceProperty   = "service_a_gateway"
	DescriptionProperty = "Http requests received counter"

	errorPropertyKey      = "error"
	statusCodePropertyKey = "status"
)

type Metrics struct {
	ErrorCount metric.CounterVec
}

func Initialize() *Metrics {
	return &Metrics{
		ErrorCount: metric.NewCounter(errorCountMetricProperties()),
	}
}

func errorCountMetricProperties() metric.Properties {
	return metric.Properties{
		Name:        NameProperty,
		Namespace:   NamespaceProperty,
		Description: DescriptionProperty,
		Type:        metric.CounterVecType,
		Properties:  []string{errorPropertyKey, statusCodePropertyKey},
	}
}
