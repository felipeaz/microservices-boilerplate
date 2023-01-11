package metric

import "github.com/prometheus/client_golang/prometheus"

const (
	CounterVecType   = "counter_vec"
	HistogramVecType = "histogram_vec"
)

type Properties struct {
	ID          string
	Name        string
	Namespace   string
	Description string
	Type        string
	Properties  []string
}

func NewCounter(p Properties) CounterVec {
	return newCounterVec(
		prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
			p.Properties,
		),
	)
}

func NewHistogram(p Properties) HistogramVec {
	return newHistogramVec(
		prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
			p.Properties,
		),
	)
}
