package hbase

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=RegionServer,sub=Memory"
func (c *Collect) parseRegionServerMemory(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"blockedFlushGauge", "unblockedFlushGauge",
			"BlockedFlushes_num_ops", "UnblockedFlushes_num_ops",
			"memStoreSize", "blockCacheSize",
			"IncreaseBlockCacheSize_num_ops", "DecreaseBlockCacheSize_num_ops",
			"IncreaseMemStoreSize_num_ops", "DecreaseMemStoreSize_num_ops",
			"tunerDoNothingCounter",
			"aboveHeapOccupancyLowWaterMarkCounter":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "regionserver_memory", metricsName),
					strings.Join([]string{c.Namespace, "regionserver memory", describeName}, " "),
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
