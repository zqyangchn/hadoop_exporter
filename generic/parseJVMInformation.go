package generic

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "java.lang:type=Threading"
func (c *PublicCollect) ParseThreading(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "PeakThreadCount", "DaemonThreadCount", "ThreadCount", "TotalStartedThreadCount":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "threading", metricsName),
					strings.Join([]string{c.Namespace, "threading", describeName}, " "),
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
func (c *PublicCollect) ParseOperatingSystem(ch chan<- prometheus.Metric, b interface{}) {
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
					prometheus.BuildFQName(c.Namespace, "os", metricsName),
					strings.Join([]string{c.Namespace, "os", describeName}, " "),
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
func (c *PublicCollect) ParseMemoryPoolCodeCache(ch chan<- prometheus.Metric, b interface{}) {
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
							prometheus.BuildFQName(c.Namespace, "memory_pool_code_cache", metricsName),
							strings.Join([]string{c.Namespace, "memory pool code cache", describeName}, " "),
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
func (c *PublicCollect) ParseMemoryPoolMetaspace(ch chan<- prometheus.Metric, b interface{}) {
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
							prometheus.BuildFQName(c.Namespace, "memory_pool_metaspace", metricsName),
							strings.Join([]string{c.Namespace, "memory pool metaspace", describeName}, " "),
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
func (c *PublicCollect) ParseMemoryPoolCompressedClassSpace(ch chan<- prometheus.Metric, b interface{}) {
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
							prometheus.BuildFQName(c.Namespace, "memory_pool_compressed_class_space", metricsName),
							strings.Join([]string{c.Namespace, "memory pool compressed class space", describeName}, " "),
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
func (c *PublicCollect) ParseBufferPoolDirect(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "TotalCapacity", "MemoryUsed", "Count":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "buffer_pool_direct", metricsName),
					strings.Join([]string{c.Namespace, "buffer pool direct", describeName}, " "),
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
func (c *PublicCollect) ParseBufferPoolMapped(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "TotalCapacity", "MemoryUsed", "Count":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "buffer_pool_mapped", metricsName),
					strings.Join([]string{c.Namespace, "hbase buffer pool mapped", describeName}, " "),
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
func (c *PublicCollect) ParseClassLoading(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "LoadedClassCount", "UnloadedClassCount", "TotalLoadedClassCount":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "class_loading", metricsName),
					strings.Join([]string{c.Namespace, "class loading", describeName}, " "),
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
