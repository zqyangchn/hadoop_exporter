package hadoop

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/zqyangchn/hadoop_exporter/common"
)

func (c *Collect) CollectMetrics() error {
	log.Debug("Start CollectMetrics ...")

	var wg sync.WaitGroup

	CollectStream := make(chan prometheus.Metric)
	defer close(CollectStream)

	StopCollectStream := make(chan struct{})
	defer close(StopCollectStream)

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		set := make([]prometheus.Metric, 0, 0)
		for {
			select {
			case m := <-CollectStream:
				set = append(set, m)
			case <-StopCollectStream:
				c.Lock()
				c.CollectMetricsSets = set
				c.Unlock()
				log.Debug("Collect Metrics Information", zap.Int("Count", len(set)))
				return
			}
		}

	}()

	req, err := http.NewRequest("GET", c.Uri, nil)
	if err != nil {
		return err
	}

	resp, err := c.hc.Do(req)
	if resp != nil {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Warn(err.Error())
			}
		}()
	}
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Warn("解析 Hadoop Json 失败", zap.Error(err))
		return err
	}

	beans := result["beans"].([]interface{})

	if common.AssertInterfaceIsNil(beans) {
		return errors.New("interface beans is nil")
	}

	for _, b := range beans {
		if common.AssertInterfaceIsNil(b) {
			log.Warn("interface b is nil")
			continue
		}

		if common.AssertInterfaceIsNil(b.(map[string]interface{})["name"]) {
			log.Warn("interface b.(map[string]interface{})[\"name\"] is nil")
			continue
		}

		switch b.(map[string]interface{})["name"] {
		// 公共
		case "java.lang:type=Threading":
			c.parseThreading(CollectStream, b)
		case "java.lang:type=OperatingSystem":
			c.parseOperatingSystem(CollectStream, b)
		case "java.lang:type=MemoryPool,name=Code Cache":
			c.parseMemoryPoolCodeCache(CollectStream, b)
		case "java.lang:type=MemoryPool,name=Metaspace":
			c.parseMemoryPoolMetaspace(CollectStream, b)
		case "java.lang:type=MemoryPool,name=Compressed Class Space":
			c.parseMemoryPoolCompressedClassSpace(CollectStream, b)
		case "java.nio:type=BufferPool,name=direct":
			c.parseBufferPoolDirect(CollectStream, b)
		case "java.nio:type=BufferPool,name=mapped":
			c.parseBufferPoolMapped(CollectStream, b)
		case "java.lang:type=ClassLoading":
			c.parseClassLoading(CollectStream, b)
		case "Hadoop:service=NameNode,name=JvmMetrics",
			"Hadoop:service=DataNode,name=JvmMetrics":
			c.parseJvmMetrics(CollectStream, b)
		case "Hadoop:service=NameNode,name=MetricsSystem,sub=Stats",
			"Hadoop:service=DataNode,name=MetricsSystem,sub=Stats":
			c.parseMetricsSystemStats(CollectStream, b)
		case "Hadoop:service=NameNode,name=UgiMetrics",
			"Hadoop:service=DataNode,name=UgiMetrics":
			c.parseUgiMetrics(CollectStream, b)
		case "Hadoop:service=NameNode,name=RpcActivityForPort" + c.namenodeServiceRPCPort,
			"Hadoop:service=DataNode,name=RpcActivityForPort" + c.datanodeRpcPort:
			c.parseRpcActivityForServiceRPCPort(CollectStream, b)

		// NameNode
		case "Hadoop:service=NameNode,name=StartupProgress":
			c.parseNameNodeStartupProgress(CollectStream, b)
		case "Hadoop:service=NameNode,name=FSNamesystem":
			c.parseNameNodeFSNamesystem(CollectStream, b)
		case "Hadoop:service=NameNode,name=RpcDetailedActivityForPort" + c.namenodeHDFSPort:
			c.parseNameNodeRpcDetailedActivityForHDFSPort(CollectStream, b)
		case "Hadoop:service=NameNode,name=RpcActivityForPort" + c.namenodeHDFSPort:
			c.parseNameNodeRpcActivityForHDFSPort(CollectStream, b)
		case "Hadoop:service=NameNode,name=RpcDetailedActivityForPort" + c.namenodeServiceRPCPort:
			c.parseNameNodeRpcDetailedActivityForServiceRPCPort(CollectStream, b)
		case "Hadoop:service=NameNode,name=NameNodeStatus":
			c.parseNameNodeNameNodeStatus(CollectStream, b)
		case "Hadoop:service=NameNode,name=NameNodeActivity":
			c.parseNameNodeActivity(CollectStream, b)
		case "Hadoop:service=NameNode,name=RetryCache.NameNodeRetryCache":
			c.parseNameNodeRetryCache(CollectStream, b)

		// DataNode
		case "Hadoop:service=DataNode,name=FSDatasetState":
			c.parseDataNodeFSDatasetState(CollectStream, b)
		case strings.Join(
			[]string{"Hadoop:service=DataNode,name=DataNodeActivity", c.Hostname, c.datanodeDataPort},
			"-"):
			c.parseDataNodeActivity(CollectStream, b)
		case "Hadoop:service=DataNode,name=RpcDetailedActivityForPort" + c.datanodeRpcPort:
			c.parseDataNodeRpcDetailedActivity(CollectStream, b)
		}
	}

	StopCollectStream <- struct{}{}
	wg.Wait()

	log.Debug("CollectMetrics Completed ...")
	return nil
}

func (c *Collect) CollectMetricsBackGround() {
	go func() {
		if err := c.CollectMetrics(); err != nil {
			panic(err)
		}

		for range time.Tick(c.CollectInterval) {
			if err := c.CollectMetrics(); err != nil {
				log.Warn("收集 Hadoop Metrics 失败", zap.Error(err))
			}
		}
	}()
}
