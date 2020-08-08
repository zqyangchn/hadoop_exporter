package hbase

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=Master,sub=FileSystem"
func (c *Collect) parseHbaseMasterFileSystem(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"HlogSplitTime_num_ops", "HlogSplitSize_num_ops",
			"MetaHlogSplitTime_num_ops", "MetaHlogSplitSize_num_ops":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "master_file_system", metricsName),
					strings.Join([]string{c.Namespace, "master file system", describeName}, " "),
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
