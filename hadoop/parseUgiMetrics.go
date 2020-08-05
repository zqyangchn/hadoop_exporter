package hadoop

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=NameNode,name=UgiMetrics"
// "Hadoop:service=DataNode,name=UgiMetrics"
func (c *Collect) parseUgiMetrics(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"LoginSuccessNumOps", "LoginSuccessAvgTime",
			"LoginFailureNumOps", "LoginFailureAvgTime",
			"GetGroupsNumOps", "GetGroupsAvgTime",
			"RenewalFailuresTotal", "RenewalFailures":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "ugi_metrics", metricsName),
					"hadoop ugi metrics "+describeName,
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
