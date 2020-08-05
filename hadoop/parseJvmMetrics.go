package hadoop

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=NameNode,name=JvmMetrics"
// "Hadoop:service=DataNode,name=JvmMetrics"
func (c *Collect) parseJvmMetrics(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"MemNonHeapUsedM", "MemNonHeapCommittedM", "MemNonHeapMaxM",
			"MemHeapUsedM", "MemHeapCommittedM", "MemHeapMaxM", "MemMaxM",
			"GcCount", "GcTimeMillis",
			"GcNumWarnThresholdExceeded", "GcNumInfoThresholdExceeded", "GcTotalExtraSleepTime",
			"ThreadsNew", "ThreadsRunnable", "ThreadsBlocked", "ThreadsWaiting", "ThreadsTimedWaiting", "ThreadsTerminated",
			"LogFatal", "LogError", "LogWarn", "LogInfo":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "jvm_metrics", metricsName),
					"hadoop jvm metrics "+describeName,
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
