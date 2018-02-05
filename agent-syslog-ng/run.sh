#!/usr/bin/env bash

export CONTAINER_NAME="agent-syslog-ng"

docker stop ${CONTAINER_NAME} && docker rm ${CONTAINER_NAME}

#docker run -d \
docker run -ti \
    --name ${CONTAINER_NAME} \
    -v "$PWD/syslog-ng.conf":/etc/syslog-ng/syslog-ng.conf \
    -v "$PWD/log":/tmp/ \
    -p 514:514 \
    -p 601:601 \
    balabit/syslog-ng:3.7


#sudo docker run -it --volumes-from [containerID for apache2] -v "$PWD/syslog-ng.conf":/etc/syslog-ng/syslog-ng.conf balabit/syslog-ng:latest