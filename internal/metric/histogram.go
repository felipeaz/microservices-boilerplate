package metric

import "github.com/prometheus/client_golang/prometheus"

type HistogramVec interface {
	Observe(duration float64, labels ...string)
}

func newHistogramVec(metric *prometheus.HistogramVec) HistogramVec {
	return &histogramVec{
		metric: metric,
	}
}

type histogramVec struct {
	metric *prometheus.HistogramVec
}

func (h *histogramVec) Observe(duration float64, labels ...string) {
	h.metric.WithLabelValues(labels...).Observe(duration)
}
