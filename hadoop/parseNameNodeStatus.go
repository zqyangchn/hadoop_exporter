package hadoop

import "github.com/prometheus/client_golang/prometheus"

// "Hadoop:service=NameNode,name=NameNodeStatus"
func (c *Collect) parseNameNodeStatus(ch chan<- prometheus.Metric, b interface{}) {
	beans := b.(map[string]interface{})

	for key, value := range beans {
		switch key {
		case "State":
			var v float64
			switch value {
			case "active":
				v = 1
			case "standby":
				v = 0
			default:
				v = 2
			}
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc(
					prometheus.BuildFQName(c.Namespace, "namenode", "status"),
					"hadoop namenode status. 0 --> standby, 1 --> active, 2 --> unknown",
					[]string{"role", "host"},
					nil,
				),
				prometheus.GaugeValue,
				v,
				c.Role,
				c.Hostname,
			)
		}
	}
}
