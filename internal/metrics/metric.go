package metrics

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

type Metric struct {
	ID          string
	Name        string
	Description string
	Type        string
	Properties  []string
}
