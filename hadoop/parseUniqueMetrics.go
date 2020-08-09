package hadoop

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

func (c *Collect) ParseUniqueMetrics(CollectStream chan prometheus.Metric, b interface{}) {
	switch b.(map[string]interface{})["name"] {
	// 公共
	case "Hadoop:service=NameNode,name=JvmMetrics",
		"Hadoop:service=DataNode,name=JvmMetrics":
		c.ParseJvmMetrics(CollectStream, b)
	case "Hadoop:service=NameNode,name=MetricsSystem,sub=Stats",
		"Hadoop:service=DataNode,name=MetricsSystem,sub=Stats":
		c.ParseMetricsSystemStats(CollectStream, b)
	case "Hadoop:service=NameNode,name=UgiMetrics",
		"Hadoop:service=DataNode,name=UgiMetrics":
		c.ParseUgiMetrics(CollectStream, b)

	// NameNode
	case "Hadoop:service=NameNode,name=StartupProgress":
		c.parseNameNodeStartupProgress(CollectStream, b)
	case "Hadoop:service=NameNode,name=FSNamesystem":
		c.parseNameNodeFSNamesystem(CollectStream, b)
	case "Hadoop:service=NameNode,name=NameNodeStatus":
		c.parseNameNodeStatus(CollectStream, b)
	case "Hadoop:service=NameNode,name=NameNodeActivity":
		c.parseNameNodeActivity(CollectStream, b)
	case "Hadoop:service=NameNode,name=RetryCache.NameNodeRetryCache":
		c.parseNameNodeRetryCache(CollectStream, b)
		// service rpc port cdh default 8022
	case "Hadoop:service=NameNode,name=RpcActivityForPort" + c.rpcActivityForPortServicePort:
		c.parseRpcActivityForServiceRPCPort(CollectStream, b)
	case "Hadoop:service=NameNode,name=RpcDetailedActivityForPort" + c.rpcActivityForPortServicePort:
		c.parseRpcDetailedActivityForServiceRPCPort(CollectStream, b)
		// hdfs port cdh default 8020
	case "Hadoop:service=NameNode,name=RpcActivityForPort" + c.rpcActivityForPortDataPort:
		c.parseRpcActivityForHDFSPort(CollectStream, b)
	case "Hadoop:service=NameNode,name=RpcDetailedActivityForPort" + c.rpcActivityForPortDataPort:
		c.parseRpcDetailedActivityForHDFSPort(CollectStream, b)

	// DataNode
	case "Hadoop:service=DataNode,name=FSDatasetState":
		c.parseDataNodeFSDatasetState(CollectStream, b)
		// RpcPort cdh default 9867
	case "Hadoop:service=DataNode,name=RpcActivityForPort" + c.rpcActivityForPortServicePort:
		c.parseRpcActivityForServiceRPCPort(CollectStream, b)
	case "Hadoop:service=DataNode,name=RpcDetailedActivityForPort" + c.rpcActivityForPortServicePort:
		c.parseDataNodeRpcDetailedActivity(CollectStream, b)
		// DataPort cdh default 9866
	case strings.Join(
		[]string{"Hadoop:service=DataNode,name=DataNodeActivity", c.Hostname, c.rpcActivityForPortDataPort},
		"-"):
		c.parseDataNodeActivity(CollectStream, b)

	default:
		/*
			log.Debug(
				"JMX Metrics Not collected completely！",
				zap.String("MetricsName", b.(map[string]interface{})["name"].(string)))

		*/
	}
}
