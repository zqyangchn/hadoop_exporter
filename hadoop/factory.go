package hadoop

import (
	"encoding/json"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/zqyangchn/hadoop_exporter/generic"
)

var log *zap.Logger

func New(role, uri string, collectMetricsBackGround bool, zapLog *zap.Logger) *Collect {
	c := new(Collect)
	c.CollectGenericMetricsForPrometheus = *generic.New(role, uri, "hadoop", zapLog)

	// init hadoop port
	switch c.Role {
	case "NameNode":
		c.InitNameNodeRPCPort()
	case "DataNode":
		c.InitDataNodeRPCPort()
	default:
		panic("hadoop_exporter -Role NameNode|DataNode")
	}

	if collectMetricsBackGround {
		c.CollectMetricsBackGround()
	}

	c.ParseMetrics = c

	log = zapLog

	return c
}

func (c *Collect) InitNameNodeRPCPort() {
	// http://beta-devicegateway-node-01.morefun-internal.com:9870/jmx
	// http://beta-devicegateway-node-01.morefun-internal.com:9870/jmx?qry=Hadoop:service=NameNode,name=NameNodeStatus

	nameNodeStatusURI := c.Uri + "?qry=Hadoop:service=NameNode,name=NameNodeStatus"
	nameNodeStatusRequest, err := http.NewRequest("GET", nameNodeStatusURI, nil)
	if err != nil {
		panic(err)
	}

	nameNodeStatusResponse, err := c.HC.Do(nameNodeStatusRequest)
	if nameNodeStatusResponse != nil {
		defer func() {
			if err := nameNodeStatusResponse.Body.Close(); err != nil {
				log.Warn(err.Error())
			}
		}()
	}
	if err != nil {
		panic(err)
	}

	var nameNodeStatusResult map[string]interface{}
	if err := json.NewDecoder(nameNodeStatusResponse.Body).Decode(&nameNodeStatusResult); err != nil {
		panic(err)
	}

	c.rpcActivityForPortDataPort = strings.Split(
		nameNodeStatusResult["beans"].([]interface{})[0].(map[string]interface{})["HostAndPort"].(string),
		":",
	)[1]

	req, err := http.NewRequest("GET", c.Uri, nil)
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

	for _, b := range result["beans"].([]interface{}) {
		beanName := b.(map[string]interface{})["name"].(string)

		if strings.Contains(beanName, "Hadoop:service=NameNode,name=RpcActivityForPort") {
			port := strings.TrimPrefix(beanName, "Hadoop:service=NameNode,name=RpcActivityForPort")

			if c.rpcActivityForPortDataPort != port {
				c.rpcActivityForPortServicePort = port
				return
			}
		}
	}
}

func (c *Collect) InitDataNodeRPCPort() {
	// http://beta-devicegateway-node-01.morefun-internal.com:9864/jmx?qry=Hadoop:service=DataNode,name=DataNodeInfo

	uri := c.Uri + "?qry=Hadoop:service=DataNode,name=DataNodeInfo"
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

	b := result["beans"].([]interface{})[0].(map[string]interface{})

	c.rpcActivityForPortDataPort = b["DataPort"].(string)
	c.rpcActivityForPortServicePort = b["RpcPort"].(string)
}
