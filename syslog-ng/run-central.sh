#!/usr/bin/env bash

docker run -d \
    --name=central \
    --net host \
    --volume $(pwd)/syslog-ng/central/syslog-ng.conf:/etc/syslog-ng/syslog-ng.conf \
    --volume $(pwd)/logs/:/tmp/ \
    -p 603:603 \
    balabit/syslog-ng:3.7