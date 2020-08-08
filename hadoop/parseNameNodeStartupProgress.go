package hadoop

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=NameNode,name=StartupProgress"
func (c *Collect) parseNameNodeStartupProgress(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"ElapsedTime", "PercentComplete",
			"LoadingFsImageCount", "LoadingFsImageElapsedTime",
			"LoadingFsImageTotal", "LoadingFsImagePercentComplete",
			"LoadingEditsCount", "LoadingEditsElapsedTime",
			"LoadingEditsTotal", "SafeModeCount",
			"SafeModeElapsedTime", "SafeModeTotal", "SafeModePercentComplete":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "namenode_startup_progress", metricsName),
					strings.Join([]string{c.Namespace, "namenode startup progress", describeName}, " "),
					[]string{
						"role",
						"host",
					},
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
