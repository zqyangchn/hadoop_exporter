package hbase

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/zqyangchn/hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=Master,sub=AssignmentManager"
func (c *Collect) parseHbaseMasterAssignmentManager(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"operationCount",
			"ritOldestAge",
			"RitDuration_num_ops",
			"ritCountOverThreshold",
			"ritCount":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "master_assignment_manager", metricsName),
					strings.Join([]string{c.Namespace, "master assignment manager", describeName}, " "),
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
