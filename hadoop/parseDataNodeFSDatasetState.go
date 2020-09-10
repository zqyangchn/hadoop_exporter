package hadoop

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"hadoop_exporter/common"
)

// "Hadoop:service=DataNode,name=FSDatasetState"
func (c *Collect) parseDataNodeFSDatasetState(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "Capacity", "DfsUsed", "Remaining",
			"NumFailedVolumes", "LastVolumeFailureDate", "EstimatedCapacityLostTotal",
			"CacheUsed", "CacheCapacity",
			"NumBlocksCached", "NumBlocksFailedToCache", "NumBlocksFailedToUnCache":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "datanode_fs_dataset_state", metricsName),
					strings.Join([]string{c.Namespace, "datanode fs dataset state", describeName}, " "),
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
