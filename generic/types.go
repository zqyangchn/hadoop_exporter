package generic

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/zqyangchn/hadoop_exporter/prom"
)

type PublicCollect struct {
	prom.PrometheusCollect

	Role     string
	Hostname string

	Namespace string

	Uri string
	HC  *http.Client

	CollectInterval    time.Duration

	Logger *zap.Logger
}
