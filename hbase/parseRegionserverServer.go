package hbase

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=RegionServer,sub=Server"
func (c *Collect) parseHbaseRegionserverServer(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			// WAL
			"hlogFileCount", "hlogFileSize",
			// 存储
			"percentFilesLocal", "storeFileSize", "storeFileIndexSize", "numReferenceFiles", "staticIndexSize", "staticBloomSize",
			// block cache
			"blockCacheFreeSize", "blockCacheCount", "blockCacheSize", "blockCacheHitCount", "blockCacheMissCount",
			"blockCacheCountHitPercent", "blockCacheExpressHitPercent",
			"blockCacheEvictionCount", "blockCacheFailedInsertionCount", "blockCacheDataMissCount",
			"blockCacheLeafIndexMissCount", "blockCacheBloomChunkMissCount", "blockCacheMetaMissCount",
			"l1CacheHitCount", "l1CacheMissCount", "l1CacheHitRatio", "l1CacheMissRatio",
			"l2CacheHitCount", "l2CacheMissCount", "l2CacheHitRatio", "l2CacheMissRatio",
			// compaction
			"compactionQueueLength", "smallCompactionQueueLength", "largeCompactionQueueLength",
			"CompactionTime_num_ops", "MajorCompactionTime_num_ops",
			"compactedCellsCount", "compactedCellsSize", "compactedOutputBytes", "compactedInputBytes",
			"majorCompactedCellsCount", "majorCompactedCellsSize", "majorCompactedInputBytes", "majorCompactedOutputBytes",
			// 请求
			"totalRequestCount", "totalRowActionRequestCount", "readRequestCount", "filteredReadRequestCount", "writeRequestCount", "blockedRequestCount",
			"rpcGetRequestCount", "rpcScanRequestCount", "rpcMultiRequestCount", "rpcMutateRequestCount",
			"slowPutCount", "slowAppendCount", "slowDeleteCount", "slowIncrementCount", "slowGetCount",
			"Bulkload_count", "Put_num_ops", "PutBatch_num_ops", "CheckAndPut_num_ops", "Increment_num_ops",
			"Delete_num_ops", "DeleteBatch_num_ops", "CheckAndDelete_num_ops", "Get_num_ops", "ScanTime_num_ops", "Replay_num_ops", "Append_num_ops",
			// 暂停
			"updatesBlockedTime", "PauseTimeWithGc_num_ops", "PauseTimeWithoutGc_num_ops", "pauseWarnThresholdExceeded", "pauseInfoThresholdExceeded",
			// region split
			"regionCount", "averageRegionSize", "splitRequestCount", "splitQueueLength", "splitSuccessCount",
			// memStore flush
			"storeCount", "storeFileCount", "memStoreSize",
			"flushQueueLength", "FlushTime_num_ops", "flushedCellsCount", "flushedCellsSize",
			"flushedOutputBytes", "flushedMemstoreBytes", "FlushOutputSize_num_ops", "FlushMemstoreSize_num_ops":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "regionserver_server", metricsName),
					strings.Join([]string{c.Namespace, "regionserver server", describeName}, " "),
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
