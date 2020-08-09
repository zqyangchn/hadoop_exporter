package generic

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

func (c *CollectGenericMetricsForPrometheus) CollectMetricsBackGround(p ParseUniqueMetrics) {
	go func() {
		if err := c.CollectMetrics(p); err != nil {
			panic(err)
		}

		for range time.Tick(c.CollectInterval) {
			if err := c.CollectMetrics(p); err != nil {
				log.Warn("Collect Metrics Failed. ", zap.Error(err))
			}
		}
	}()
}

func (c *CollectGenericMetricsForPrometheus) CollectMetrics(p ParseUniqueMetrics) error {
	log.Debug("Start CollectMetrics ...")

	var wg sync.WaitGroup

	CollectStream := make(chan prometheus.Metric)
	defer close(CollectStream)

	StopCollectStream := make(chan struct{})
	defer close(StopCollectStream)

	// update CollectMetricsSets
	wg.Add(1)
	go func() {
		defer wg.Done()
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
		p.ParseExporterStatus(CollectStream, err)
		return err
	}

	resp, err := c.HC.Do(req)
	if resp != nil {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Warn("resp.Body.Close() failed !", zap.Error(err))
			}
		}()
	}
	if err != nil {
		p.ParseExporterStatus(CollectStream, err)
		return err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		p.ParseExporterStatus(CollectStream, err)
		return err
	}

	beans := result["beans"].([]interface{})
	if common.AssertInterfaceIsNil(beans) {
		err := errors.New("interface beans is nil")
		p.ParseExporterStatus(CollectStream, err)
		return err
	}

	// parse this exporter status
	p.ParseExporterStatus(CollectStream, nil)

	c.ParseGenericMetrics(CollectStream, beans, p)

	StopCollectStream <- struct{}{}
	wg.Wait()

	log.Debug("CollectMetrics Completed ...")

	return nil
}

func (c *CollectGenericMetricsForPrometheus) GetPrometheusMetrics() []prometheus.Metric {
	defer c.Unlock()

	c.Lock()
	return c.CollectMetricsSets
}
