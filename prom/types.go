package prom

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

// Implement Prometheus Interface Collect Struct
type PrometheusCollect struct {
	sync.Mutex

	CollectMetricsSets []prometheus.Metric
}