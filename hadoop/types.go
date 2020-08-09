package hadoop

import (
	"github.com/zqyangchn/hadoop_exporter/generic"
)

type Collect struct {
	generic.CollectGenericMetricsForPrometheus

	// hadoop namenode datanode
	namenodeHDFSPort       string
	namenodeServiceRPCPort string

	datanodeRpcPort  string
	datanodeDataPort string
}
