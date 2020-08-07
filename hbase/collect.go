package hbase

import (
	"encoding/json"
	"errors"
	"net/http"
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
		c.parseHBaseStatus(CollectStream, err)
		return err
	}

	resp, err := c.hc.Do(req)
	if resp != nil {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Warn("resp.Body.Close() failed !", zap.Error(err))
			}
		}()
	}
	if err != nil {
		c.parseHBaseStatus(CollectStream, err)
		return err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.parseHBaseStatus(CollectStream, err)
		return err
	}

	beans := result["beans"].([]interface{})
	if common.AssertInterfaceIsNil(beans) {
		err := errors.New("interface beans is nil")
		c.parseHBaseStatus(CollectStream, err)
		return err
	}

	c.parseHBaseStatus(CollectStream, nil)

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
		case "Hadoop:service=HBase,name=JvmMetrics":
			c.parseJvmMetrics(CollectStream, b)
		case "Hadoop:service=HBase,name=MetricsSystem,sub=Stats":
			c.parseMetricsSystemStats(CollectStream, b)
		case "Hadoop:service=HBase,name=UgiMetrics":
			c.parseUgiMetrics(CollectStream, b)

		// HMaster
		case "Hadoop:service=HBase,name=Master,sub=Balancer":
			c.parseHbaseMasterBalancer(CollectStream, b)
		case "Hadoop:service=HBase,name=Master,sub=AssignmentManager":
			c.parseHbaseMasterAssignmentManager(CollectStream, b)
		case "Hadoop:service=HBase,name=Master,sub=FileSystem":
			c.parseHbaseMasterFileSystem(CollectStream, b)
		case "Hadoop:service=HBase,name=Master,sub=Server":
			c.parseHbaseMasterServer(CollectStream, b)
		case "Hadoop:service=HBase,name=Master,sub=IPC":
			c.parseHbaseMasterIPC(CollectStream, b)

		// RegionServer
		case "Hadoop:service=HBase,name=RegionServer,sub=Regions":
			c.parseHbaseRegionServerRegions(CollectStream, b)
		case "Hadoop:service=HBase,name=RegionServer,sub=IO":
			c.parseHbaseRegionServerIO(CollectStream, b)
		case "Hadoop:service=HBase,name=RegionServer,sub=TableLatencies":
		//	c.parseHbaseRegionServerTableLatencies(CollectStream, b)
		case "Hadoop:service=HBase,name=RegionServer,sub=WAL":
			c.parseHbaseRegionServerWAL(CollectStream, b)
		case "Hadoop:service=HBase,name=RegionServer,sub=Tables":
			c.parseHbaseRegionServerTables(CollectStream, b)
		case "Hadoop:service=HBase,name=RegionServer,sub=Server":
			c.parseHbaseRegionserverServer(CollectStream, b)
		case "Hadoop:service=HBase,name=RegionServer,sub=IPC":
			c.parseHbaseRegionserverIPC(CollectStream, b)
		case "Hadoop:service=HBase,name=RegionServer,sub=Memory":

		case "Hadoop:service=HBase,name=RegionServer,sub=Replication":

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
