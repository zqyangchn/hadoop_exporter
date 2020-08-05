package hadoop

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=NameNode,name=NameNodeActivity"
func (c *Collect) parseNameNodeActivity(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"FilesCreated", "FilesAppended", "FilesRenamed", "FilesTruncated", "FilesDeleted", "FilesInGetListingOps",

			"TotalFileOps", "CreateFileOps", "DeleteFileOps",

			"GetBlockLocations",

			"GetListingOps", "FileInfoOps", "AddBlockOps", "GetAdditionalDatanodeOps", "CreateSymlinkOps",
			"GetLinkTargetOps", "AllowSnapshotOps", "DisallowSnapshotOps", "CreateSnapshotOps", "DeleteSnapshotOps",
			"RenameSnapshotOps", "ListSnapshottableDirOps", "SnapshotDiffReportOps", "BlockReceivedAndDeletedOps",

			"SuccessfulReReplications",
			"NumTimesReReplicationNotScheduled",
			"TimeoutReReplications",

			"BlockOpsQueued",
			"BlockOpsBatched",

			"SafeModeTime",
			"FsImageLoadTime",
			"TransactionsNumOps", "TransactionsAvgTime", "TransactionsBatchedInSync",

			"SyncsNumOps", "SyncsAvgTime",
			"StorageBlockReportNumOps", "StorageBlockReportAvgTime",
			"CacheReportNumOps", "CacheReportAvgTime",
			"GenerateEDEKTimeNumOps", "GenerateEDEKTimeAvgTime",
			"WarmUpEDEKTimeNumOps", "WarmUpEDEKTimeAvgTime",
			"ResourceCheckTimeNumOps", "ResourceCheckTimeAvgTime",

			"EditLogTailTimeNumOps", "EditLogTailTimeAvgTime",
			"EditLogFetchTimeNumOps", "EditLogFetchTimeAvgTime",
			"NumEditLogLoadedNumOps", "NumEditLogLoadedAvgCount",
			"EditLogTailIntervalNumOps", "EditLogTailIntervalAvgTime",
			"GetEditNumOps", "GetEditAvgTime",
			"GetImageNumOps", "GetImageAvgTime",
			"PutImageNumOps", "PutImageAvgTime":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_activity", metricsName),
					"hadoop namenode activity "+describeName,
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
