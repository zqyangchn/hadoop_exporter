package hadoop

import (
	"hadoop_exporter/generic"
)

type Collect struct {
	generic.CollectGenericMetricsForPrometheus

	rpcActivityForPortDataPort    string
	rpcActivityForPortServicePort string
}
