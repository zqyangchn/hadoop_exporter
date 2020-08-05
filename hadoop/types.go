package hadoop

import (
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "hadoop"
)

// Hadoop Collect 结构体
type Collect struct {
	sync.Mutex

	Role     string
	Hostname string

	Uri string
	hc  *http.Client

	namenodeHDFSPort       string
	namenodeServiceRPCPort string

	datanodeRpcPort  string
	datanodeDataPort string

	CollectInterval    time.Duration
	CollectMetricsSets []prometheus.Metric
}
