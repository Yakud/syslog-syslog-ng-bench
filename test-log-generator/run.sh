#!/usr/bin/env bash

export IMAGE_NAME="yakud/test-gen-log"
export CONTAINER_NAME="log-gen"

docker stop ${CONTAINER_NAME} && docker rm ${CONTAINER_NAME}

docker build -t ${IMAGE_NAME} .

#docker run -ti \
docker run -d \
    --name ${CONTAINER_NAME} \
    --log-driver syslog \
    --log-opt syslog-address=tcp://127.0.0.1:601 \
    --log-opt mode=non-blocking \
    --log-opt max-buffer-size=400m \
    --log-opt tag="{{.ImageName}}|{{.Name}}|{{.ID}}" \
    ${IMAGE_NAME}

