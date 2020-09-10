package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"hadoop_exporter/common"
	"hadoop_exporter/hadoop"
)

const (
	// namenode default
	defaultRole = "NameNode"
	defaultURI  = "http://beta-devicegateway-node-01.morefun-internal.com:9870"

	// datanode default
	//defaultRole = "DataNode"
	//defaultURI  = "http://beta-devicegateway-node-01.morefun-internal.com:9864"
)

var (
	log *zap.Logger

	LogLevel      string
	LogOutput     string
	LogOutputFile string

	Role string
	Uri  string

	ListenUri string
)

func init() {
	// 处理命令行参数解析
	flag.StringVar(&LogLevel, "LogLevel", "debug", "日志级别: debug|info")
	flag.StringVar(&LogOutput, "LogOutput", "stdout", "日志输出: stdout|file")
	flag.StringVar(&LogOutputFile, "LogOutputFile", "/var/log/hadoop_exporter/hadoop_exporter.log", "日志文件")

	flag.StringVar(&Role, "Role", defaultRole, "NameNode|DataNode")
	flag.StringVar(&Uri, "Uri", defaultURI, "监控节点 URI")

	flag.StringVar(&ListenUri, "ListenUri", "0.0.0.0:18428", "web 监听 ip:port")

	flag.Parse()

	// 初始日志
	log = common.Initialization(LogLevel, LogOutput, LogOutputFile)

	// 注册 prometheus 监控 go-client 路由
	http.Handle("/metrics", promhttp.Handler())

	// 注册 prometheus Collector 接口结构
	prometheus.MustRegister(
		hadoop.New(Role, Uri, true, log),
	)

	log.Info(
		"Prometheus Monitor Hadoop JMX Metrics",
		zap.String("Role", Role),
		zap.String("Uri", Uri),
	)
}

func main() {
	server := &http.Server{
		Addr:              ListenUri,
		ReadTimeout:       1 * time.Minute,
		ReadHeaderTimeout: 20 * time.Second,
		WriteTimeout:      2 * time.Minute,
	}

	log.Info("Web Starting Completed !", zap.String("ListenUri", ListenUri))
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
