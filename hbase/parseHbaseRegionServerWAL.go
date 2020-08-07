package hbase

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=RegionServer,sub=WAL"
func (c *Collect) parseHbaseRegionServerWAL(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"appendCount", "slowAppendCount", "AppendSize_num_ops", "AppendTime_num_ops",
			"SyncTime_num_ops",
			"rollRequest",
			"writtenBytes",
			"lowReplicaRollRequest":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "regionserver_wal", metricsName),
					"hbase regionserver wal "+describeName,
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
