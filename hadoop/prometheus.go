package hadoop

import "github.com/prometheus/client_golang/prometheus"

func (c *Collect) Collect(ch chan<- prometheus.Metric) {
	c.Lock()
	defer c.Unlock()

	for _, m := range c.CollectMetricsSets {
		ch <- m
	}
}

func (c *Collect) Describe(ch chan<- *prometheus.Desc) {
	c.Lock()
	defer c.Unlock()

	for _, m := range c.CollectMetricsSets {
		ch <- m.Desc()
	}
}
