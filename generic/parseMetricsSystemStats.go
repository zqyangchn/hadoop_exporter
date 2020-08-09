package generic

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=MetricsSystem,sub=Stats"
func (c *CollectGenericMetricsForPrometheus) ParseMetricsSystemStats(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"NumActiveSources", "NumAllSources",
			"NumActiveSinks", "NumAllSinks",
			"SnapshotNumOps", "SnapshotAvgTime",
			"PublishNumOps", "PublishAvgTime", "DroppedPubAll":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "metrics_system_stats", metricsName),
					strings.Join([]string{c.Namespace, "metrics system stats", describeName}, " "),
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		}
	}
}
