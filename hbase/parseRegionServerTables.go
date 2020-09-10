package hbase

import (
	"go.uber.org/zap"
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=RegionServer,sub=Tables"
func (c *Collect) parseHbaseRegionServerTables(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})
	re := regexp.MustCompile(`^Namespace_([\w-]+)_table_([\w-]+)_metric_(.*)$`)

	for key, value := range beans {
		switch {
		case key == "numTables":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "regionserver_tables", metricsName),
					strings.Join([]string{c.Namespace, "regionserver tables", describeName}, " "),
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role, c.Hostname,
			)
		case strings.Contains(key, "Namespace_"):
			match := re.FindStringSubmatch(key)
			nameSpace, tableName, metrics := match[1], match[2], match[3]

			switch metrics {
			case
				"totalRequestCount",
				"writeRequestCount",
				"readRequestCount",
				"storeFileSize",
				"tableSize",
				"memstoreSize":
				metricsName, describeName := common.ConversionToPrometheusFormat(metrics)
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc(
						prometheus.BuildFQName(c.Namespace, "regionserver_tables", metricsName),
						strings.Join([]string{c.Namespace, "regionserver tables", describeName}, " "),
						[]string{"role", "host", "namespace", "table"},
						nil,
					),
					prometheus.GaugeValue,
					value.(float64),
					c.Role, c.Hostname, nameSpace, tableName,
				)
			default:
				log.Debug("parseHbaseRegionServerTables incomplete indicator collection",
					zap.String("metrics", metrics),
					zap.String("nameSpace", nameSpace),
					zap.String("tableName", tableName),
				)
			}
		}
	}
}
