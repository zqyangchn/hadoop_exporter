package hbase

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

// Generate hbase_node_status 0: dead 1: active
func (c *Collect) parseExporterStatus(ch chan<- prometheus.Metric, err error) {
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(c.Namespace, "node", "status"),
				"hbase node status 0: dead 1: active",
				[]string{"role", "host"},
				nil,
			),
			prometheus.GaugeValue,
			0,
			c.Role,
			c.Hostname,
		)
		log.Warn("set hbase_node_status 0", zap.Error(err))
		return
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(c.Namespace, "node", "status"),
			"hbase node status 0: dead 1: active",
			[]string{"role", "host"},
			nil,
		),
		prometheus.GaugeValue,
		1,
		c.Role,
		c.Hostname,
	)
}
