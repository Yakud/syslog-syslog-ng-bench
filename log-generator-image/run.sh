#!/usr/bin/env bash

docker run -d \
    --net host \
    --log-driver syslog \
    --log-opt syslog-address="tcp://127.0.0.1:603" \
    --log-opt tag="{{.ImageName}}|{{.Name}}|{{.ID}}" \
    log-generator-image
