package hadoop

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"hadoop_exporter/common"
)

func (c *Collect) parseRpcActivityForServiceRPCPort(ch chan<- prometheus.Metric, b interface{}) {
	// "Hadoop:service=NameNode,name=RpcActivityForPort8022"
	// "Hadoop:service=DataNode,name=RpcActivityForPort9867"

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
					prometheus.BuildFQName(c.Namespace, "rpc_activity_for_service_rpc_port", metricsName),
					strings.Join([]string{c.Namespace, "rpc activity for service rpc port", describeName}, " "),
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

func (c *Collect) parseRpcDetailedActivityForServiceRPCPort(ch chan<- prometheus.Metric, b interface{}) {
	// "Hadoop:service=NameNode,name=RpcDetailedActivityForPort8022"

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
					prometheus.BuildFQName(c.Namespace,
						"rpc_detailed_activity_for_service_rpc_port", metricsName),
					strings.Join([]string{c.Namespace, "rpc detailed activity for service rpc port", describeName}, " "),
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

func (c *Collect) parseRpcActivityForHDFSPort(ch chan<- prometheus.Metric, b interface{}) {
	// "Hadoop:service=NameNode,name=RpcActivityForPort8020"

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
					prometheus.BuildFQName(c.Namespace, "rpc_activity_for_hdfs_port", metricsName),
					strings.Join([]string{c.Namespace, "rpc activity for hdfs port", describeName}, " "),
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

func (c *Collect) parseRpcDetailedActivityForHDFSPort(ch chan<- prometheus.Metric, b interface{}) {
	// "Hadoop:service=NameNode,name=RpcDetailedActivityForPort8020"

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
					prometheus.BuildFQName(c.Namespace,
						"rpc_detailed_activity_for_hdfs_port", metricsName),
					strings.Join([]string{c.Namespace, "rpc detailed activity for hdfs port", describeName}, " "),
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
