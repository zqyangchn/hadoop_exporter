package hadoop

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=NameNode,name=FSNamesystem"
func (c *Collect) parseNameNodeFSNamesystem(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"TotalLoad", "TotalSyncCount", "FilesTotal", "VolumeFailuresTotal",
			"CapacityTotal", "CapacityUsed", "CapacityRemaining", "CapacityUsedNonDFS",
			"BlockCapacity", "EstimatedCapacityLostTotal",
			"BlocksTotal", "ExcessBlocks", "CorruptBlocks", "MissingBlocks", "LowRedundancyBlocks",
			"TotalReplicatedBlocks", "UnderReplicatedBlocks", "PostponedMisreplicatedBlocks",
			"CorruptReplicatedBlocks", "MissingReplOneBlocks", "MissingReplicatedBlocks",
			"MissingReplicationOneBlocks", "BytesInFutureReplicatedBlocks", "PendingDeletionReplicatedBlocks",
			"LowRedundancyReplicatedBlocks", "HighestPriorityLowRedundancyReplicatedBlocks",
			"PendingDeletionBlocks", "PendingReplicationBlocks", "PendingReconstructionBlocks",
			"ScheduledReplicationBlocks",
			"ExpiredHeartbeats", "TransactionsSinceLastCheckpoint", "TransactionsSinceLastLogRoll",
			"LastWrittenTransactionId", "LastCheckpointTime",
			"LowRedundancyECBlockGroups", "CorruptECBlockGroups", "MissingECBlockGroups",
			"BytesInFutureECBlockGroups", "PendingDeletionECBlocks", "HighestPriorityLowRedundancyECBlocks",
			"TotalECBlockGroups",
			"Snapshots", "SnapshottableDirectories",
			"LockQueueLength", "NumActiveClients",
			"NumEncryptionZones", "NumFilesUnderConstruction", "NumTimedOutPendingReconstructions",
			"PendingDataNodeMessageCount", "MillisSinceLastLoadedEdits",
			"StaleDataNodes", "NumLiveDataNodes", "NumDeadDataNodes", "NumDecomLiveDataNodes",
			"NumDecomDeadDataNodes", "NumDecommissioningDataNodes", "NumInMaintenanceLiveDataNodes",
			"NumInMaintenanceDeadDataNodes", "NumEnteringMaintenanceDataNodes",
			"NumStaleStorages":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_fs_name_system", metricsName),
					"hadoop namenode fs name system "+describeName,
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
