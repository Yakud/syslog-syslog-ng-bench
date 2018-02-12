#!/usr/bin/env bash

    --net host \
docker run -d \
    --log-driver syslog \
    --log-opt syslog-address="tcp://127.0.0.1:603" \
    --log-opt tag="{{.ImageName}}|{{.Name}}|{{.ID}}" \
    log-generator-image
