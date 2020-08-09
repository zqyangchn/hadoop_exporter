package generic

import (
	"net"
	"net/http"
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

	log = zapLog

	return c
}
