# Hadoop 集群相关的 Prometheus Exporter

##### Test on Hadoop 3.0.0-cdh6.2.0 HBase 2.1.0-cdh6.2.0 ####

### 编译Linux平台的 Exporter
	./build.sh

### hadoop_exporter  查看启动命令行参数
	./hadoop_exporter -h

### Grafana  展示
	grafana_example 目录包含 grafana 的 json 文件(HadoopMetricsFromJMX.json)
	grafana 模板不包含所有监控指标, 但 hadoop_exporter 暴露了, 自行添加即可
