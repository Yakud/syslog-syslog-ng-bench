#!/usr/bin/env bash

docker run -d \
    --name=log-generator \
    --log-driver syslog \
    --log-opt syslog-address="tcp://0.0.0.0:603" \
    --log-opt tag="{{.ImageName}}|{{.Name}}|{{.ID}}" \
    log-generator-image
