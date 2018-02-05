#!/usr/bin/env bash

env IMAGE_NAME="agent-syslog-ng"

docker run -it \
    -v "$PWD/syslog-ng.conf":/etc/syslog-ng/syslog-ng.conf \
     balabit/syslog-ng:latest