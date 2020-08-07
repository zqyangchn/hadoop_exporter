package hbase

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=Master,sub=IPC"
func (c *Collect) parseHbaseMasterIPC(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			// bytes
			"sentBytes", "receivedBytes",
			// connection and handler
			"numOpenConnections", "numLifoModeSwitches", "numActiveHandler",
			"numActiveWriteHandler", "numActiveReadHandler", "numActiveScanHandler",
			// queue
			"queueSize", "numGeneralCallsDropped",
			"numCallsInGeneralQueue", "numCallsInReplicationQueue", "numCallsInPriorityQueue",
			"numCallsInReadQueue", "numCallsInScanQueue", "numCallsInWriteQueue",
			//
			"TotalCallTime_num_ops", "ProcessCallTime_num_ops", "QueueCallTime_num_ops",
			"RequestSize_num_ops", "ResponseSize_num_ops",
			// auth
			"authorizationSuccesses", "authorizationFailures",
			"authenticationSuccesses", "authenticationFailures", "authenticationFallbacks":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "master_ipc", metricsName),
					"hbase master ipc "+describeName,
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case
			"exceptions",
			"exceptions.multiResponseTooLarge",
			"exceptions.callQueueTooBig":

			key = strings.Replace(key, ".", "_", -1)

			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "master_ipc", metricsName),
					"hbase master ipc "+describeName,
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case
			"exceptions.RegionMovedException",
			"exceptions.UnknownScannerException",
			"exceptions.OutOfOrderScannerNextException",
			"exceptions.RegionTooBusyException",
			"exceptions.FailedSanityCheckException",
			"exceptions.NotServingRegionException",
			"exceptions.ScannerResetException":

			key = strings.Replace(key, ".", "", -1)

			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(namespace, "master_ipc", metricsName),
					"hbase master ipc "+describeName,
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
