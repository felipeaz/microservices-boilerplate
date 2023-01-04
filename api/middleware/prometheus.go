package middleware

import (
	"github.com/gin-gonic/gin"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"strconv"
)

const (
	defaultMetricsPath = "/metrics"
)

type commonMetrics struct {
	requestCounterMetric  *prom.CounterVec
	requestDurationMetric *prom.HistogramVec
}

type prometheus struct {
	router      *gin.Engine
	metricsPath string
	metrics     commonMetrics
}

func NewPrometheusMiddleware(router *gin.Engine) Middleware {
	p := &prometheus{
		router:      router,
		metricsPath: defaultMetricsPath,
		metrics:     buildStandardMetrics(),
	}
	p.registerRoutePath()
	return p
}

func (p *prometheus) HandleFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == p.metricsPath {
			c.Next()
			return
		}

		method := c.Request.Method
		host := c.Request.Host
		status := strconv.Itoa(c.Writer.Status())

		timer := prom.NewTimer(p.metrics.requestDurationMetric.WithLabelValues(status, method, path))
		timer.ObserveDuration()

		p.metrics.requestCounterMetric.WithLabelValues(
			status,
			method,
			c.HandlerName(),
			host,
			path,
		).Inc()
	}
}

func (p *prometheus) register() {
	p.registerRoutePath()
	p.registerMetrics()
}

func (p *prometheus) registerRoutePath() {
	p.router.GET(p.metricsPath, gin.WrapH(promhttp.Handler()))
}

func (p *prometheus) registerMetrics() {
	err := prom.Register(p.metrics.requestCounterMetric)
	if err != nil {
		log.Println("failed to register standard counter metric")
	}
	err = prom.Register(p.metrics.requestDurationMetric)
	if err != nil {
		log.Println("failed to register standard duration metric")
	}
}

func buildStandardMetrics() commonMetrics {
	return commonMetrics{
		requestCounterMetric: prom.NewCounterVec(
			prom.CounterOpts{
				Name: "http_request_count",
				Help: "Number of http request",
			},
			[]string{"code", "method", "handler", "host", "url"},
		),
		requestDurationMetric: prom.NewHistogramVec(
			prom.HistogramOpts{
				Name: "http_request_latency_in_sec",
				Help: "Http request latency in sec",
			},
			[]string{"code", "method", "url"},
		),
	}
}
