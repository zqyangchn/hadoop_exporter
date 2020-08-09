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
	case "Hadoop:service=NameNode,name=RpcActivityForPort" + c.namenodeServiceRPCPort,
		"Hadoop:service=DataNode,name=RpcActivityForPort" + c.datanodeRpcPort:
		c.parseRpcActivityForServiceRPCPort(CollectStream, b)

	// NameNode
	case "Hadoop:service=NameNode,name=StartupProgress":
		c.parseNameNodeStartupProgress(CollectStream, b)
	case "Hadoop:service=NameNode,name=FSNamesystem":
		c.parseNameNodeFSNamesystem(CollectStream, b)
	case "Hadoop:service=NameNode,name=RpcDetailedActivityForPort" + c.namenodeHDFSPort:
		c.parseNameNodeRpcDetailedActivityForHDFSPort(CollectStream, b)
	case "Hadoop:service=NameNode,name=RpcActivityForPort" + c.namenodeHDFSPort:
		c.parseNameNodeRpcActivityForHDFSPort(CollectStream, b)
	case "Hadoop:service=NameNode,name=RpcDetailedActivityForPort" + c.namenodeServiceRPCPort:
		c.parseNameNodeRpcDetailedActivityForServiceRPCPort(CollectStream, b)
	case "Hadoop:service=NameNode,name=NameNodeStatus":
		c.parseNameNodeStatus(CollectStream, b)
	case "Hadoop:service=NameNode,name=NameNodeActivity":
		c.parseNameNodeActivity(CollectStream, b)
	case "Hadoop:service=NameNode,name=RetryCache.NameNodeRetryCache":
		c.parseNameNodeRetryCache(CollectStream, b)

	// DataNode
	case "Hadoop:service=DataNode,name=FSDatasetState":
		c.parseDataNodeFSDatasetState(CollectStream, b)
	case strings.Join(
		[]string{"Hadoop:service=DataNode,name=DataNodeActivity", c.Hostname, c.datanodeDataPort},
		"-"):
		c.parseDataNodeActivity(CollectStream, b)
	case "Hadoop:service=DataNode,name=RpcDetailedActivityForPort" + c.datanodeRpcPort:
		c.parseDataNodeRpcDetailedActivity(CollectStream, b)
	}

}
