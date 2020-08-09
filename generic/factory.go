package generic

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

var log *zap.Logger

func New(role, uri, namespace string, zapLog *zap.Logger) *CollectGenericMetricsForPrometheus {
	c := new(CollectGenericMetricsForPrometheus)

	c.Namespace = namespace
	c.Role = role
	c.Uri = uri + "/jmx"

	c.CollectInterval = 5 * time.Second

	c.HC = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,

			MaxIdleConns:          10,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
		},
		Timeout: 2 * time.Second,
	}

	c.initHostname()

	log = zapLog

	return c
}

func (c *CollectGenericMetricsForPrometheus) initHostname() {
	/*
		Get Hostname by jmx api
		hadoop
		http://beta-devicegateway-node-01.morefun-internal.com:9864/jmx?qry=java.lang:type=Runtime
		http://beta-devicegateway-node-01.morefun-internal.com:9870/jmx?qry=java.lang:type=Runtime
		hbase
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
