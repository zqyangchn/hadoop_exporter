package generic

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

type CollectGenericMetricsForPrometheus struct {
	sync.Mutex

	Role     string
	Hostname string

	Namespace string

	Uri string
	HC  *http.Client

	CollectInterval    time.Duration
	CollectMetricsSets []prometheus.Metric

	Logger *zap.Logger
}

type ParseUniqueMetrics interface {
	ParseExporterStatus (ch chan<- prometheus.Metric, err error)
	ParseUniqueMetrics (chan prometheus.Metric, interface{})
}