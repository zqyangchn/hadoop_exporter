package hbase

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=RegionServer,sub=IO"
func (c *Collect) parseHbaseRegionServerIO(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"fsChecksumFailureCount",
			"FsWriteTime_num_ops", "FsReadTime_num_ops":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "regionserver_io", metricsName),
					"hbase regionserver_io "+describeName,
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "FsPReadTime_num_ops":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "regionserver_io", "fs_pre_read_time_num_ops"),
					"hbase regionserver io fs pre-read time num ops",
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
