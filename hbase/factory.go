package hbase

import (
	"go.uber.org/zap"

	"hadoop_exporter/generic"
)

var log *zap.Logger

func New(role, uri string, collectMetricsBackGround bool, zapLog *zap.Logger) *Collect {
	c := new(Collect)
	c.CollectGenericMetricsForPrometheus = *generic.New(role, uri, "hbase", zapLog)

	if collectMetricsBackGround {
		c.CollectMetricsBackGround()
	}

	c.ParseMetrics = c

	log = zapLog

	return c
}
