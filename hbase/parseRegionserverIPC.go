package hbase

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=RegionServer,sub=IPC"
func (c *Collect) parseHbaseRegionserverIPC(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			// connection and handler
			"numOpenConnections", "numActiveHandler", "numActiveWriteHandler", "numActiveReadHandler", "numActiveScanHandler",
			// IPC queue
			"queueSize", "numCallsInGeneralQueue", "numCallsInReplicationQueue", "numCallsInPriorityQueue",
			"numCallsInWriteQueue", "numCallsInReadQueue", "numCallsInScanQueue",
			// IPC call count
			"TotalCallTime_num_ops", "RequestSize_num_ops", "ResponseSize_num_ops", "ProcessCallTime_num_ops", "QueueCallTime_num_ops",
			// IPC bytes
			"sentBytes", "receivedBytes":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "regionserver_ipc", metricsName),
					strings.Join([]string{c.Namespace, "regionserver ipc", describeName}, " "),
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case
			// IPC exception count
			"exceptions",
			"exceptions.callQueueTooBig",
			"exceptions.multiResponseTooLarge":
			key = strings.Replace(key, ".", "_", -1)

			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "regionserver_ipc", metricsName),
					strings.Join([]string{c.Namespace, "regionserver ipc", describeName}, " "),
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
			"exceptions.RegionTooBusyException",
			"exceptions.FailedSanityCheckException",
			"exceptions.UnknownScannerException",
			"exceptions.OutOfOrderScannerNextException",
			"exceptions.ScannerResetException",
			"exceptions.NotServingRegionException":
			key = strings.Replace(key, ".", "", -1)

			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "regionserver_ipc", metricsName),
					strings.Join([]string{c.Namespace, "regionserver ipc", describeName}, " "),
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
