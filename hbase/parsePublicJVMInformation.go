package hbase

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "java.lang:type=Threading"
func (c *Collect) parseThreading(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "PeakThreadCount", "DaemonThreadCount", "ThreadCount", "TotalStartedThreadCount":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "threading", metricsName),
					"hbase threading "+describeName,
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

// "java.lang:type=OperatingSystem"
func (c *Collect) parseOperatingSystem(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"OpenFileDescriptorCount", "MaxFileDescriptorCount",
			"FreePhysicalMemorySize", "TotalPhysicalMemorySize",
			"SystemCpuLoad", "ProcessCpuLoad", "SystemLoadAverage",
			"AvailableProcessors":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "os", metricsName),
					"hbase os "+describeName,
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
							"hbase memory pool code cache "+describeName,
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
							"hbase memory pool metaspace "+describeName,
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
							"hbase memory pool compressed class space "+describeName,
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

// "java.nio:type=BufferPool,name=direct"
func (c *Collect) parseBufferPoolDirect(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "TotalCapacity", "MemoryUsed", "Count":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "buffer_pool_direct", metricsName),
					"hbase buffer pool direct "+describeName,
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

// "java.nio:type=BufferPool,name=mapped"
func (c *Collect) parseBufferPoolMapped(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "TotalCapacity", "MemoryUsed", "Count":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "buffer_pool_mapped", metricsName),
					"hbase buffer pool mapped "+describeName,
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

// "java.lang:type=ClassLoading"
func (c *Collect) parseClassLoading(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "LoadedClassCount", "UnloadedClassCount", "TotalLoadedClassCount":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "class_loading", metricsName),
					"hbase class loading "+describeName,
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
