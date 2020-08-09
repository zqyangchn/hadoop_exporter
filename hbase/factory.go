package hbase

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zqyangchn/hadoop_exporter/generic"
	"go.uber.org/zap"
)

var log *zap.Logger

func New(role, uri string, collectMetricsBackGround bool, zapLog *zap.Logger) *Collect {
	c := new(Collect)
	c.CollectGenericMetricsForPrometheus = generic.New(role, uri, "hbase", zapLog)

	c.initInformation()

	if collectMetricsBackGround {
		c.CollectMetricsBackGround(c)
	}

	log = zapLog

	return c
}

func (c *Collect) initInformation() {
	/*
		Get Hostname by jmx api
		http://beta-devicegateway-node-01.morefun-internal.com:16010/jmx?qry=java.lang:type=Runtime
		http://beta-devicegateway-node-01.morefun-internal.com:16030/jmx?qry=java.lang:type=Runtime
	*/
	uri := c.Uri + "?qry=java.lang:type=Runtime"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}

	resp, err := c.HC.Do(req)
	if resp != nil {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Warn(err.Error())
			}
		}()
	}
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	c.Hostname = strings.Split(
		result["beans"].([]interface{})[0].(map[string]interface{})["Name"].(string),
		"@",
	)[1]
}
