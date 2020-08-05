#!/bin/bash

# Mac
#go build -gcflags "-m" -ldflags "-w" hadoop_exporter.go

# Linux
GOOS=linux GOARCH=amd64 go build -gcflags "-m" -ldflags "-w" hadoop_exporter.go
