package hadoop

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=NameNode,name=MetricsSystem,sub=Stats"
// "Hadoop:service=DataNode,name=MetricsSystem,sub=Stats"
func (c *Collect) parseMetricsSystemStats(ch chan<- prometheus.Metric, b interface{}) {
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
					prometheus.BuildFQName(namespace, "metrics_system_stats", metricsName),
					"hadoop metrics system stats "+describeName,
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
