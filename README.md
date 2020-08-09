# Hadoop 集群相关的 Prometheus Exporter

##### Test on Hadoop 3.0.0-cdh6.2.0 HBase 2.1.0-cdh6.2.0 ####

### 编译Mac或Linux平台的 Exporter
	./build.sh mac
	./build.sh linux

### hadoop_exporter  查看启动命令行参数
######	hadoop_exporter -h
	-ListenUri string
		web 监听 ip:port (default "0.0.0.0:18428")
	-LogLevel string
		日志级别: debug|info (default "debug")
	-LogOutput string
		日志输出: stdout|file (default "stdout")
	-LogOutputFile string
		日志文件 (default "/var/log/hadoop_exporter/hadoop_exporter.log")
	-Role string
		NameNode|DataNode (default "NameNode")
	-Uri string
		监控节点 URI (default "http://beta-devicegateway-node-01.morefun-internal.com:9870")


### hbase_exporter  查看启动命令行参数
###### hbase_exporter -h
	-ListenUri string
		web 监听 ip:port (default "0.0.0.0:8428")
	-LogLevel string
		日志级别: debug|info (default "debug")
	-LogOutput string
		日志输出: stdout|file (default "stdout")
	-LogOutputFile string
		日志文件 (default "/var/log/hbase_exporter/hbase_exporter.log")
	-Role string
		HMaster|RegionServer (default "RegionServer")
	-Uri string
		监控节点 URI (default "http://beta-devicegateway-node-01.morefun-internal.com:16030")
### Grafana  展示
	.example-grafana 目录含 grafana 的 json 文件(HadoopMetricsFromJMX.json)
	模板没有画出所有监控指标, hadoop_exporter 和 hbase_exporter 输出的指标可以自行添加
