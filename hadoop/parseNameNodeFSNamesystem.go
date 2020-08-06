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
			"CapacityTotal", "CapacityUsed", "CapacityRemaining",
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
		case "CapacityUsedNonDFS":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_fs_name_system", "capacity_used_non_dfs"),
					"hadoop namenode fs name system capacity used non dfs",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "LowRedundancyECBlockGroups":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_fs_name_system", "low_redundancy_ec_block_groups"),
					"hadoop namenode fs name system low redundancy ec block groups",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "CorruptECBlockGroups":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_fs_name_system", "corrupt_ec_block_groups"),
					"hadoop namenode fs name system corrupt ec block groups",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "MissingECBlockGroups":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_fs_name_system", "missing_ec_block_groups"),
					"hadoop namenode fs name system missing ec block groups",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "BytesInFutureECBlockGroups":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_fs_name_system", "bytes_in_future_ec_block_groups"),
					"hadoop namenode fs name system bytes in future ec block groups",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "PendingDeletionECBlocks":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_fs_name_system", "pending_deletion_ec_blocks"),
					"hadoop namenode fs name system pending deletion ec blocks",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "HighestPriorityLowRedundancyECBlocks":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_fs_name_system", "highest_priority_low_redundancy_ec_blocks"),
					"hadoop namenode fs name system highest priority low redundancy ec blocks",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "TotalECBlockGroups":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_fs_name_system", "total_ec_block_groups"),
					"hadoop namenode fs name system total ec block groups",
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
