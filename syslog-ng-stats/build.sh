#!/usr/bin/env bash

export GOROOT=/usr/local/go/
export GOPATH=$(pwd)/syslog-ng-stats/src/

env GOOS=linux GOARCH=amd64 go build -o bin/syslog-ng-stats syslog-ng-stats/src//main.go
go build -o bin/syslog-ng-stats-osx syslog-ng-stats/src/main.go