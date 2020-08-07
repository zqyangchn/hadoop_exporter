#!/bin/bash

Usage(){
    echo "Usage: 
            ./build mac|linux
            Compile For Mac or Linux
        "
    exit 2
}

if [ $# -ne 1 ]; then
    Usage
fi

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct

if [ "$1" = "mac" ]; then
    echo "Compile For Mac ..."
    go build -gcflags "-m" -ldflags "-w" -o hadoop_exporter hadoop-exporter/hadoop_exporter.go && \
        go build -gcflags "-m" -ldflags "-w" -o hbase_exporter hbase-exporter/hbase_exporter.go
    echo 
    echo " Compiled For Mac Done !"
elif [ "$1" = "linux" ]; then
    echo "Compile For Linux ..."
    GOOS=linux GOARCH=amd64 go build -gcflags "-m" -ldflags "-w" -o hadoop_exporter hadoop-exporter/hadoop_exporter.go && \
        GOOS=linux GOARCH=amd64 go build -gcflags "-m" -ldflags "-w" -o hbase_exporter hbase-exporter/hbase_exporter.go
    echo 
    echo " Compiled For Linux Done !"
elif [ "$1" = "rm" ]; then
    rm -f hadoop_exporter hbase_exporter
else
    Usage
fi

