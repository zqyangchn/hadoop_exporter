package hadoop

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zqyangchn/hadoop_exporter/generic"
	"go.uber.org/zap"
)

var log *zap.Logger

func New(role, uri, namenodeHDFSPort, namenodeServiceRPCPort string, collectMetricsBackGround bool, zapLog *zap.Logger) *Collect {
	c := new(Collect)
	c.CollectGenericMetricsForPrometheus = generic.New(role, uri, "hadoop", zapLog)

	c.namenodeHDFSPort = namenodeHDFSPort
	c.namenodeServiceRPCPort = namenodeServiceRPCPort

	switch c.Role {
	case "NameNode":
		c.InitNameNodeInformation()
	case "DataNode":
		c.InitDataNodeInformation()
	default:
		panic("Unknown role, Enter the correct role. Usage ./hadoop_exporter -role roleName")
	}

	if collectMetricsBackGround {
		c.CollectMetricsBackGround(c)
	}

	log = zapLog

	return c
}

// http://beta-devicegateway-node-01.morefun-internal.com:9870/jmx?qry=Hadoop:service=NameNode,name=NameNodeStatus
func (c *Collect) InitNameNodeInformation() {
	uri := c.Uri + "?qry=Hadoop:service=NameNode,name=NameNodeStatus"
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

	hostAndPort := result["beans"].([]interface{})[0].(map[string]interface{})["HostAndPort"].(string)

	c.Hostname = strings.Split(hostAndPort, ":")[0]
	c.namenodeHDFSPort = strings.Split(hostAndPort, ":")[1]
}

// http://beta-devicegateway-node-01.morefun-internal.com:9864/jmx?qry=Hadoop:service=DataNode,name=DataNodeInfo
func (c *Collect) InitDataNodeInformation() {
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

	c.Hostname = b["DatanodeHostname"].(string)
	c.datanodeRpcPort = b["RpcPort"].(string)
	c.datanodeDataPort = b["DataPort"].(string)
}
