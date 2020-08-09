package hbase

import (
	"go.uber.org/zap"

	"github.com/zqyangchn/hadoop_exporter/generic"
)

var log *zap.Logger

func New(role, uri string, collectMetricsBackGround bool, zapLog *zap.Logger) *Collect {
	c := new(Collect)
	c.CollectGenericMetricsForPrometheus = generic.New(role, uri, "hbase", zapLog)

	if collectMetricsBackGround {
		c.CollectMetricsBackGround(c)
	}

	log = zapLog

	return c
}
