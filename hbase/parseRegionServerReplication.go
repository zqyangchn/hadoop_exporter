package hbase

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=RegionServer,sub=Replication"
func (c *Collect) parseRegionServerReplication(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"sink.appliedOps",
			"sink.appliedBatches",
			"Sink.ageOfLastAppliedOp_num_ops":
			key = strings.Replace(key, ".", "_", -1)

			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "regionserver_replication", metricsName),
					strings.Join([]string{c.Namespace, "regionserver replication", describeName}, " "),
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "sink.appliedHFiles":
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "regionserver_replication",
						"sink_applied_hfiles"),
					strings.Join([]string{c.Namespace, "regionserver replication", "sink applied hfiles"}, " "),
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
