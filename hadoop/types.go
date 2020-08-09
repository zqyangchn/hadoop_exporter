package hadoop

import (
	"github.com/zqyangchn/hadoop_exporter/generic"
)

type Collect struct {
	generic.CollectGenericMetricsForPrometheus

	rpcActivityForPortDataPort    string
	rpcActivityForPortServicePort string
}
