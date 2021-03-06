package hadoop

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"hadoop_exporter/common"
)

// "Hadoop:service=DataNode,name=DataNodeActivity-beta-devicegateway-node-01.morefun-internal.com-9866"
func (c *Collect) parseDataNodeActivity(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"BytesWritten", "TotalWriteTime", "BytesRead", "TotalReadTime",
			"BlocksWritten", "BlocksRead", "BlocksReplicated", "BlocksRemoved",
			"BlocksVerified", "BlockVerificationFailures", "BlocksCached", "BlocksUncached",
			"ReadsFromLocalClient", "ReadsFromRemoteClient",
			"WritesFromLocalClient", "WritesFromRemoteClient",
			"BlocksGetLocalPathInfo", "RemoteBytesRead", "RemoteBytesWritten",
			"RamDiskBlocksWrite", "RamDiskBlocksWriteFallback", "RamDiskBytesWrite",
			"RamDiskBlocksReadHits", "RamDiskBlocksEvicted", "RamDiskBlocksEvictedWithoutRead",
			"RamDiskBlocksEvictionWindowMsNumOps", "RamDiskBlocksEvictionWindowMsAvgTime",
			"RamDiskBlocksLazyPersisted", "RamDiskBlocksDeletedBeforeLazyPersisted",
			"RamDiskBytesLazyPersisted", "RamDiskBlocksLazyPersistWindowMsNumOps",
			"RamDiskBlocksLazyPersistWindowMsAvgTime",
			"FsyncCount", "VolumeFailures",
			"DatanodeNetworkErrors", "DataNodeActiveXceiversCount",
			"ReadBlockOpNumOps", "ReadBlockOpAvgTime",
			"WriteBlockOpNumOps", "WriteBlockOpAvgTime",
			"BlockChecksumOpNumOps", "BlockChecksumOpAvgTime",
			"CopyBlockOpNumOps", "CopyBlockOpAvgTime",
			"ReplaceBlockOpNumOps", "ReplaceBlockOpAvgTime",
			"HeartbeatsNumOps", "HeartbeatsAvgTime",
			"HeartbeatsTotalNumOps", "HeartbeatsTotalAvgTime",
			"LifelinesNumOps", "LifelinesAvgTime",
			"BlockReportsNumOps", "BlockReportsAvgTime",
			"IncrementalBlockReportsNumOps", "IncrementalBlockReportsAvgTime",
			"CacheReportsNumOps", "CacheReportsAvgTime",
			"PacketAckRoundTripTimeNanosNumOps", "PacketAckRoundTripTimeNanosAvgTime",
			"FlushNanosNumOps", "FlushNanosAvgTime", "FsyncNanosNumOps", "FsyncNanosAvgTime",
			"SendDataPacketBlockedOnNetworkNanosNumOps", "SendDataPacketBlockedOnNetworkNanosAvgTime",
			"SendDataPacketTransferNanosNumOps", "SendDataPacketTransferNanosAvgTime",
			"EcReconstructionTasks", "EcFailedReconstructionTasks",
			"EcDecodingTimeNanos", "EcReconstructionBytesRead",
			"EcReconstructionBytesWritten", "EcReconstructionRemoteBytesRead",
			"EcReconstructionReadTimeMillis", "EcReconstructionDecodingTimeMillis", "EcReconstructionWriteTimeMillis":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "datanode_activity", metricsName),
					strings.Join([]string{c.Namespace, "datanode activity", describeName}, " "),
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "BlocksInPendingIBR":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "datanode_activity", "blocks_in_pending_ibr"),
					"hadoop datanode activity blocks in pending ibr",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "BlocksReceivingInPendingIBR":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "datanode_activity", "blocks_receiving_in_pending_ibr"),
					"hadoop datanode activity blocks receiving in pending ibr",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "BlocksReceivedInPendingIBR":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "datanode_activity", "blocks_received_in_pending_ibr"),
					"hadoop datanode activity blocks received in pending ibr",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "BlocksDeletedInPendingIBR":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "datanode_activity", "blocks_deleted_in_pending_ibr"),
					"hadoop datanode activity blocks deleted in pending ibr",
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
