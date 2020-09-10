package hbase

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"

	"hadoop_exporter/common"
)

// "Hadoop:service=HBase,name=Master,sub=Balancer"
func (c *Collect) parseHbaseMasterBalancer(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "miscInvocationCount", "BalancerCluster_num_ops":
			metricsName, describeName := common.ConversionToPrometheusFormat(key)
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "master_balancer", metricsName),
					strings.Join([]string{c.Namespace, "master balancer", describeName}, " "),
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				value.(float64),
				c.Role,
				c.Hostname,
			)
		case "tag.isBalancerActive":
			if value.(string) == "true" {
				ch <- prometheus.MustNewConstMetric(
					prometheus.NewDesc(
						prometheus.BuildFQName(c.Namespace, "master_balancer", "is_active_balancer"),
						"hbase master balancer is active balancer 1: active, 0: not active",
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
					prometheus.BuildFQName(c.Namespace, "master_balancer", "is_active_balancer"),
					"hbase master balancer is active balancer 1: active, 0: not active",
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
