package hadoop

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zqyangchn/hadoop_exporter/common"
)

// "java.lang:type=MemoryPool,name=Code Cache"
func (c *Collect) parseMemoryPoolCodeCache(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "Usage":
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "init", "committed", "used", "max":
					metricsName, describeName := common.ConversionToPrometheusFormat(k)
					ch <- prometheus.MustNewConstMetric(
						prometheus.NewDesc(
							prometheus.BuildFQName(namespace, "memory_pool_code_cache", metricsName),
							"hadoop memory pool code cache "+describeName,
							[]string{"role", "host"},
							nil,
						),
						prometheus.GaugeValue,
						v.(float64),
						c.Role,
						c.Hostname,
					)
				}
			}
		}
	}
}

// "java.lang:type=MemoryPool,name=Metaspace"
func (c *Collect) parseMemoryPoolMetaspace(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "Usage":
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "init", "committed", "used", "max":
					metricsName, describeName := common.ConversionToPrometheusFormat(k)
					ch <- prometheus.MustNewConstMetric(
						prometheus.NewDesc(
							prometheus.BuildFQName(namespace, "memory_pool_metaspace", metricsName),
							"hadoop memory pool metaspace "+describeName,
							[]string{"role", "host"},
							nil,
						),
						prometheus.GaugeValue,
						v.(float64),
						c.Role,
						c.Hostname,
					)
				}
			}
		}
	}
}

// "java.lang:type=MemoryPool,name=Compressed Class Space"
func (c *Collect) parseMemoryPoolCompressedClassSpace(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "Usage":
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "init", "committed", "used", "max":
					metricsName, describeName := common.ConversionToPrometheusFormat(k)
					ch <- prometheus.MustNewConstMetric(
						prometheus.NewDesc(
							prometheus.BuildFQName(namespace, "memory_pool_compressed_class_space", metricsName),
							"hadoop memory pool compressed class space "+describeName,
							[]string{"role", "host"},
							nil,
						),
						prometheus.GaugeValue,
						v.(float64),
						c.Role,
						c.Hostname,
					)
				}
			}
		}
	}
}
