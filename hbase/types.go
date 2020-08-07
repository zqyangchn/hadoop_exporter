package hbase

import (
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "hbase"
)

// Hbase Collect 结构体
type Collect struct {
	sync.Mutex

	Role     string
	Hostname string

	Uri string
	hc  *http.Client

	CollectInterval    time.Duration
	CollectMetricsSets []prometheus.Metric
}
