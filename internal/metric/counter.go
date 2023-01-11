package metric

import "github.com/prometheus/client_golang/prometheus"

type CounterVec interface {
	Increment(labels ...string)
}

func newCounterVec(metric *prometheus.CounterVec) CounterVec {
	return &counterVec{
		metric: metric,
	}
}

type counterVec struct {
	metric *prometheus.CounterVec
}

func (c *counterVec) Increment(labels ...string) {
	c.metric.WithLabelValues(labels...).Inc()
}
