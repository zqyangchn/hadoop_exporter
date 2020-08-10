package hbase

import (
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=RegionServer,sub=Regions"
func (c *Collect) parseHbaseRegionServerRegions(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	re := regexp.MustCompile(`^Namespace_([\w-]+)_table_([\w-]+)_region_([\w-]+)_metric_([\w-]+)$`)

	for key, value := range beans {
		if strings.Contains(key, "Namespace_") {
			match := re.FindStringSubmatch(key)
			nameSpace, tableName, region, metrics := match[1], match[2], match[3], match[4]

			switch metrics {
			case
				"replicaid",
				"numReferenceFiles",
				// ops
				"getCount", "putCount", "scanCount", "appendCount", "deleteCount", "incrementCount",
				// read write
				"readRequestCount", "writeRequestCount", "filteredReadRequestCount",
				// Flush
				"maxFlushQueueSize",
				"flushesQueuedCount",
				// store
				"storeCount", "memStoreSize", "storeFileSize", "storeFileCount",
				"minStoreFileAge", "avgStoreFileAge", "maxStoreFileAge",
				// compaction
				"compactionsQueuedCount", "compactionsFailedCount", "compactionsCompletedCount",
				"maxCompactionQueueSize", "lastMajorCompactionAge",
				"numFilesCompactedCount", "numBytesCompactedCount":
				metricsName, describeName := common.ConversionToPrometheusFormat(metrics)
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc(
						prometheus.BuildFQName(c.Namespace, "regionserver_regions", metricsName),
						strings.Join([]string{c.Namespace, "regionserver regions", describeName}, " "),
						[]string{"role", "host", "namespace", "table", "region"},
						nil,
					),
					prometheus.GaugeValue,
					value.(float64),
					c.Role, c.Hostname, nameSpace, tableName, region,
				)
			default:
				log.Debug("parseHbaseRegionServerRegions incomplete indicator collection",
					zap.String("metrics", metrics))
			}
		}
	}
}
