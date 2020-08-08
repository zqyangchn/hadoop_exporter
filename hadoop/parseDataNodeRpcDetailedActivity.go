package hadoop

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=DataNode,name=RpcDetailedActivityForPort9867"
func (c *Collect) parseDataNodeRpcDetailedActivity(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"InitReplicaRecoveryNumOps", "InitReplicaRecoveryAvgTime",
			"UpdateReplicaUnderRecoveryNumOps", "UpdateReplicaUnderRecoveryAvgTime":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "rpc_detailed_activity", metricsName),
					strings.Join([]string{c.Namespace, "rpc detailed activity", describeName}, " "),
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
