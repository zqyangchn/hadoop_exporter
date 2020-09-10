package hbase

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=Master,sub=Server"
func (c *Collect) parseHbaseMasterServer(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case
			"mergePlanCount", "splitPlanCount",
			"masterActiveTime", "masterStartTime", "masterFinishedInitializationTime",
			"averageLoad",
			"numRegionServers", "numDeadRegionServers",
			"clusterRequests":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "master_server", metricsName),
					strings.Join([]string{c.Namespace, "master server", describeName}, " "),
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "tag.isActiveMaster":
			if value.(string) == "true" {
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc(
						prometheus.BuildFQName(c.Namespace, "master_server", "is_active_master"),
						"hbase master server is active master 1: active, 0: not active",
						[]string{"role", "host"},
						nil,
					),
					prometheus.GaugeValue,
					1,
					c.Role,
					c.Hostname,
				)
				continue
			}
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "master_server", "is_active_master"),
					"hbase master server is active master 1: active, 0: not active",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				0,
				c.Role,
				c.Hostname,
			)

		}
	}
}
