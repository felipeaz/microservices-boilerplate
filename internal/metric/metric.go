package metric

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
	Namespace   string
	Description string
	Type        string
	Properties  []string
}

func New(p Properties) prometheus.Collector {
	switch p.Type {
	case CounterVec:
		return prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
			p.Properties,
		)
	case Counter:
		return prometheus.NewCounter(
			prometheus.CounterOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
		)
	case GaugeVec:
		return prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
			p.Properties,
		)
	case Gauge:
		return prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
		)
	case HistogramVec:
		return prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
			p.Properties,
		)
	case Histogram:
		return prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
		)
	case SummaryVec:
		return prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
			p.Properties,
		)
	case Summary:
		return prometheus.NewSummary(
			prometheus.SummaryOpts{
				Name:      p.Name,
				Namespace: p.Namespace,
				Help:      p.Description,
			},
		)
	default:
		return nil
	}
}
