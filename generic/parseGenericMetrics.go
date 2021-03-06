package generic

import (
	"github.com/prometheus/client_golang/prometheus"

	"hadoop_exporter/common"
)

func (c *CollectGenericMetricsForPrometheus) ParseGenericMetrics(CollectStream chan prometheus.Metric, beans []interface{}) {
	for _, b := range beans {
		if common.AssertInterfaceIsNil(b) {
			log.Warn("interface b is nil")
			continue
		}

		if common.AssertInterfaceIsNil(b.(map[string]interface{})["name"]) {
			log.Warn("interface b.(map[string]interface{})[\"name\"] is nil")
			continue
		}

		switch b.(map[string]interface{})["name"] {
		// 公共
		case "java.lang:type=Threading":
			c.ParseThreading(CollectStream, b)
		case "java.lang:type=OperatingSystem":
			c.ParseOperatingSystem(CollectStream, b)
		case "java.lang:type=MemoryPool,name=Code Cache":
			c.ParseMemoryPoolCodeCache(CollectStream, b)
		case "java.lang:type=MemoryPool,name=Metaspace":
			c.ParseMemoryPoolMetaspace(CollectStream, b)
		case "java.lang:type=MemoryPool,name=Compressed Class Space":
			c.ParseMemoryPoolCompressedClassSpace(CollectStream, b)
		case "java.nio:type=BufferPool,name=direct":
			c.ParseBufferPoolDirect(CollectStream, b)
		case "java.nio:type=BufferPool,name=mapped":
			c.ParseBufferPoolMapped(CollectStream, b)
		case "java.lang:type=ClassLoading":
			c.ParseClassLoading(CollectStream, b)
		default:
			// parse Unique Metrics
			c.ParseMetrics.ParseUniqueMetrics(CollectStream, b)
		}
	}
}
