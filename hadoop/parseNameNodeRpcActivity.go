package hadoop

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=NameNode,name=RpcDetailedActivityForPort8020"
func (c *Collect) parseNameNodeRpcDetailedActivityForHDFSPort(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"GetBlockLocationsNumOps", "GetBlockLocationsAvgTime",
			"ListCachePoolsNumOps", "ListCachePoolsAvgTime",
			"FileNotFoundExceptionNumOps", "FileNotFoundExceptionAvgTime",
			"DeleteNumOps", "DeleteAvgTime",
			"GetServerDefaultsNumOps", "GetServerDefaultsAvgTime",
			"FsyncNumOps", "FsyncAvgTime",
			"AddBlockNumOps", "AddBlockAvgTime",
			"SetStoragePolicyNumOps", "SetStoragePolicyAvgTime",
			"ListEncryptionZonesNumOps", "ListEncryptionZonesAvgTime",
			"CreateNumOps", "CreateAvgTime",
			"SetPermissionNumOps", "SetPermissionAvgTime",
			"IsFileClosedNumOps", "IsFileClosedAvgTime",
			"SetSafeModeNumOps", "SetSafeModeAvgTime",
			"GetListingNumOps", "GetListingAvgTime",
			"Rename2NumOps", "Rename2AvgTime",
			"ListCacheDirectivesNumOps", "ListCacheDirectivesAvgTime",
			"AbandonBlockNumOps", "AbandonBlockAvgTime",
			"StandbyExceptionNumOps", "StandbyExceptionAvgTime",
			"RenewLeaseNumOps", "RenewLeaseAvgTime",
			"RenameNumOps", "RenameAvgTime",
			"MkdirsNumOps", "MkdirsAvgTime",
			"RecoverLeaseNumOps", "RecoverLeaseAvgTime",
			"SetTimesNumOps", "SetTimesAvgTime",
			"CompleteNumOps", "CompleteAvgTime",
			"GetFileInfoNumOps", "GetFileInfoAvgTime":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace,
						"namenode_rpc_detailed_activity_for_hdfs_port", metricsName),
					"hadoop namenode rpc detailed activity for hdfs port "+describeName,
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

// "Hadoop:service=NameNode,name=RpcActivityForPort8020"
func (c *Collect) parseNameNodeRpcActivityForHDFSPort(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"CallQueueLength",
			"ReceivedBytes", "SentBytes",
			"RpcQueueTimeNumOps", "RpcQueueTimeAvgTime",
			"RpcProcessingTimeNumOps", "RpcProcessingTimeAvgTime",
			"DeferredRpcProcessingTimeNumOps", "DeferredRpcProcessingTimeAvgTime",
			"RpcAuthenticationFailures", "RpcAuthenticationSuccesses",
			"RpcAuthorizationFailures", "RpcAuthorizationSuccesses",
			"RpcClientBackoff", "RpcSlowCalls",
			"NumOpenConnections", "NumDroppedConnections":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "namenode_rpc_activity_for_hdfs_port", metricsName),
					"hadoop namenode rpc activity for hdfs port "+describeName,
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

// "Hadoop:service=NameNode,name=RpcActivityForPort8022"
// "Hadoop:service=DataNode,name=RpcActivityForPort9867"
func (c *Collect) parseRpcActivityForServiceRPCPort(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"CallQueueLength",
			"ReceivedBytes", "SentBytes",
			"RpcQueueTimeNumOps", "RpcQueueTimeAvgTime",
			"RpcProcessingTimeNumOps", "RpcProcessingTimeAvgTime",
			"DeferredRpcProcessingTimeNumOps", "DeferredRpcProcessingTimeAvgTime",
			"RpcAuthenticationFailures", "RpcAuthenticationSuccesses",
			"RpcAuthorizationFailures", "RpcAuthorizationSuccesses",
			"RpcClientBackoff", "RpcSlowCalls",
			"NumOpenConnections", "NumDroppedConnections":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "rpc_activity_for_service_rpc_port", metricsName),
					"hadoop rpc activity for service rpc port"+describeName,
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

// "Hadoop:service=NameNode,name=RpcDetailedActivityForPort8022"
func (c *Collect) parseNameNodeRpcDetailedActivityForServiceRPCPort(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"GetServiceStatusNumOps", "GetServiceStatusAvgTime",
			"RollEditLogNumOps", "RollEditLogAvgTime",
			"RegisterDatanodeNumOps", "RegisterDatanodeAvgTime",
			"SendHeartbeatNumOps", "SendHeartbeatAvgTime",
			"CacheReportNumOps", "CacheReportAvgTime",
			"IOExceptionNumOps", "IOExceptionAvgTime",
			"MonitorHealthNumOps", "MonitorHealthAvgTime",
			"VersionRequestNumOps", "VersionRequestAvgTime",
			"CommitBlockSynchronizationNumOps", "CommitBlockSynchronizationAvgTime",
			"TransitionToActiveNumOps", "TransitionToActiveAvgTime",
			"TransitionToStandbyNumOps", "TransitionToStandbyAvgTime",
			"BlockReportNumOps", "BlockReportAvgTime",
			"BlockReceivedAndDeletedNumOps", "BlockReceivedAndDeletedAvgTime":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace,
						"namenode_rpc_detailed_activity_for_service_rpc_port", metricsName),
					"hadoop namenode rpc detailed activity for service rpc port "+describeName,
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
