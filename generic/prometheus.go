package generic

import "github.com/prometheus/client_golang/prometheus"

func (c *CollectGenericMetricsForPrometheus) Collect(ch chan<- prometheus.Metric) {
	for _, m := range c.GetPrometheusMetrics() {
		ch <- m
	}
}

func (c *CollectGenericMetricsForPrometheus) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.GetPrometheusMetrics() {
		ch <- m.Desc()
	}
}