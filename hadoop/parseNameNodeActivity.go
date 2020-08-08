package hadoop

import (
	"strings"

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
					prometheus.BuildFQName(c.Namespace, "namenode_activity", metricsName),
					strings.Join([]string{c.Namespace, "namenode activity", describeName}, " "),
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "GenerateEDEKTimeNumOps":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "namenode_activity", "generate_edek_time_num_ops"),
					"hadoop namenode activity generate edek time num ops",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "GenerateEDEKTimeAvgTime":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "namenode_activity", "generate_edek_time_avg_time"),
					"hadoop namenode activity generate edek time avg time",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "WarmUpEDEKTimeNumOps":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "namenode_activity", "warm_up_edek_time_num_ops"),
					"hadoop namenode activity warm up edek time num ops",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "WarmUpEDEKTimeAvgTime":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "namenode_activity", "warm_up_edek_time_avg_time"),
					"hadoop namenode activity warm up edek time avg time",
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
