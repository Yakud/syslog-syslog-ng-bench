#!/usr/bin/env bash

docker run -d \
    --name=log-generator \
    --net host \
    --log-driver syslog \
    --log-opt syslog-address="upd://127.0.0.1:603" \
    log-generator-image