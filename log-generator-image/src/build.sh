#!/usr/bin/env bash

export GOROOT=/usr/local/go/
export GOPATH=$(pwd)/log-generator-image/src/

env GOOS=linux GOARCH=amd64 go build -o bin/gen-log log-generator-image/src/main.go
go build -o bin/gen-log-osx log-generator-image/src/main.go