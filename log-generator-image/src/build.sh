#!/usr/bin/env bash

export GOROOT=/usr/local/go/
export GOPATH=/Users/yakud/work/analytics-system/logs/log-generator-image/src/

env GOOS=linux GOARCH=amd64 go build -o log-generator-image/gen-log log-generator-image/src/main.go
go build -o log-generator-image/gen-log-osx log-generator-image/src/main.go