package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	CounterVec   = "counter_vec"
	Counter      = "counter"
	GaugeVec     = "gauge_vec"
	Gauge        = "gauge"
	HistogramVec = "histogram_vec"
	Histogram    = "histogram"
	SummaryVec   = "summary_vec"
	Summary      = "summary"
)

type Properties struct {
	ID          string
	Name        string
	Description string
	Type        string
	Properties  []string
}

type MetricCollector interface {
	New(p Properties) prometheus.Collector
}

func NewMetricCollector() MetricCollector {
	return &metric{}
}

type metric struct{}
