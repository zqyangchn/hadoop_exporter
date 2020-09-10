package hbase

import (
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=RegionServer,sub=TableLatencies"
func (c *Collect) parseHbaseRegionServerTableLatencies(ch chan<- prometheus.Metric, b interface{}) {
	bean := b.(map[string]interface{})

	re := regexp.MustCompile(`^Namespace_([\w-]+)_table_([\w-]+)_metric_([\w-]+)$`)

	for key, value := range bean {
		if strings.Contains(key, "_num_ops") {
			match := re.FindStringSubmatch(key)
			nameSpace, tableName, metrics := match[1], match[2], match[3]

			switch metrics {
			case
				"putTime_num_ops",
				"getTime_num_ops",
				"scanSize_num_ops",
				"scanTime_num_ops",
				"appendTime_num_ops",
				"deleteTime_num_ops",
				"incrementTime_num_ops",
				"putBatchTime_num_ops",
				"deleteBatchTime_num_ops":
				metricsName, describeName := common.ConversionToPrometheusFormat(metrics)
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc(
						prometheus.BuildFQName(c.Namespace, "regionserver_table_latencies", metricsName),
						strings.Join([]string{c.Namespace, "regionserver table latencies", describeName}, " "),
						[]string{"role", "host", "namespace", "table"},
						nil,
					),
					prometheus.GaugeValue,
					value.(float64),
					c.Role, c.Hostname, nameSpace, tableName,
				)
			default:
				log.Debug("parseHbaseRegionServerTableLatencies incomplete indicator collection",
					zap.String("metrics", metrics))
			}
		}
	}
}
