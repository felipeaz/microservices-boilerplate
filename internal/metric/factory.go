package metric

import "github.com/prometheus/client_golang/prometheus"

func NewMetric(m *Metric) prometheus.Collector {
	switch m.Type {
	case CounterVec:
		return prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: m.Name,
				Help: m.Description,
			},
			m.Args,
		)
	case Counter:
		return prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: m.Name,
				Help: m.Description,
			},
		)
	case GaugeVec:
		return prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: m.Name,
				Help: m.Description,
			},
			m.Args,
		)
	case Gauge:
		return prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: m.Name,
				Help: m.Description,
			},
		)
	case HistogramVec:
		return prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: m.Name,
				Help: m.Description,
			},
			m.Args,
		)
	case Histogram:
		return prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name: m.Name,
				Help: m.Description,
			},
		)
	case SummaryVec:
		return prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name: m.Name,
				Help: m.Description,
			},
			m.Args,
		)
	case Summary:
		return prometheus.NewSummary(
			prometheus.SummaryOpts{
				Name: m.Name,
				Help: m.Description,
			},
		)
	default:
		return nil
	}
}
