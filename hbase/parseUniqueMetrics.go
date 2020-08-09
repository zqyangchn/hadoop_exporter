package hbase

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (c *Collect) ParseUniqueMetrics(CollectStream chan prometheus.Metric, b interface{}) {
	switch b.(map[string]interface{})["name"] {
	// 公共
	case "Hadoop:service=HBase,name=JvmMetrics":
		c.ParseJvmMetrics(CollectStream, b)
	case "Hadoop:service=HBase,name=MetricsSystem,sub=Stats":
		c.ParseMetricsSystemStats(CollectStream, b)
	case "Hadoop:service=HBase,name=UgiMetrics":
		c.ParseUgiMetrics(CollectStream, b)

	// HMaster
	case "Hadoop:service=HBase,name=Master,sub=Balancer":
		c.parseHbaseMasterBalancer(CollectStream, b)
	case "Hadoop:service=HBase,name=Master,sub=AssignmentManager":
		c.parseHbaseMasterAssignmentManager(CollectStream, b)
	case "Hadoop:service=HBase,name=Master,sub=FileSystem":
		c.parseHbaseMasterFileSystem(CollectStream, b)
	case "Hadoop:service=HBase,name=Master,sub=Server":
		c.parseHbaseMasterServer(CollectStream, b)
	case "Hadoop:service=HBase,name=Master,sub=IPC":
		c.parseHbaseMasterIPC(CollectStream, b)

	// RegionServer
	case "Hadoop:service=HBase,name=RegionServer,sub=Regions":
		c.parseHbaseRegionServerRegions(CollectStream, b)
	case "Hadoop:service=HBase,name=RegionServer,sub=IO":
		c.parseHbaseRegionServerIO(CollectStream, b)
	case "Hadoop:service=HBase,name=RegionServer,sub=TableLatencies":
	//	c.parseHbaseRegionServerTableLatencies(CollectStream, b)
	case "Hadoop:service=HBase,name=RegionServer,sub=WAL":
		c.parseHbaseRegionServerWAL(CollectStream, b)
	case "Hadoop:service=HBase,name=RegionServer,sub=Tables":
		c.parseHbaseRegionServerTables(CollectStream, b)
	case "Hadoop:service=HBase,name=RegionServer,sub=Server":
		c.parseHbaseRegionserverServer(CollectStream, b)
	case "Hadoop:service=HBase,name=RegionServer,sub=IPC":
		c.parseHbaseRegionserverIPC(CollectStream, b)
	case "Hadoop:service=HBase,name=RegionServer,sub=Memory":

	case "Hadoop:service=HBase,name=RegionServer,sub=Replication":

	}
}
