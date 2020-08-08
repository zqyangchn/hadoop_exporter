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

func (c *PublicCollect) CollectMetricsBackGround(
	parseExporterStatus func(ch chan<- prometheus.Metric, err error),
	parseUniqueMetrics func(chan prometheus.Metric, interface{}),
) {
	go func() {
		if c.Logger == nil {
			panic("Logger is nil, please use *zap.logger")
		}

		if err := c.CollectMetrics(parseExporterStatus, parseUniqueMetrics); err != nil {
			panic(err)
		}

		for range time.Tick(c.CollectInterval) {
			if err := c.CollectMetrics(parseExporterStatus, parseUniqueMetrics); err != nil {
				c.Logger.Warn("Collect Metrics Failed. ", zap.Error(err))
			}
		}
	}()
}

func (c *PublicCollect) CollectMetrics(
	parseExporterStatus func(ch chan<- prometheus.Metric, err error),
	parseUniqueMetrics func(chan prometheus.Metric, interface{}),
) error {
	c.Logger.Debug("Start CollectMetrics ...")

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
				if c.Logger != nil {
					c.Logger.Debug("Collect Metrics Information", zap.Int("Count", len(set)))
				}
				return
			}
		}

	}()

	req, err := http.NewRequest("GET", c.Uri, nil)
	if err != nil {
		parseExporterStatus(CollectStream, err)
		return err
	}

	resp, err := c.HC.Do(req)
	if resp != nil {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				c.Logger.Warn("resp.Body.Close() failed !", zap.Error(err))
			}
		}()
	}
	if err != nil {
		parseExporterStatus(CollectStream, err)
		return err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		parseExporterStatus(CollectStream, err)
		return err
	}

	beans := result["beans"].([]interface{})
	if common.AssertInterfaceIsNil(beans) {
		err := errors.New("interface beans is nil")
		parseExporterStatus(CollectStream, err)
		return err
	}

	// parse this exporter status
	parseExporterStatus(CollectStream, nil)

	c.ParseGenericMetrics(CollectStream, beans, parseUniqueMetrics)

	StopCollectStream <- struct{}{}
	wg.Wait()

	c.Logger.Debug("CollectMetrics Completed ...")

	return nil
}