#!/usr/bin/env bash

docker stop central && docker rm central

docker run -d \
    --net host \
    --name=central \
    --volume $(pwd)/syslog-ng/central/syslog-ng.conf:/etc/syslog-ng/syslog-ng.conf \
    --volume /home/ssd-logs/:/tmp/ \
    -p 603:603 \
    balabit/syslog-ng:3.7