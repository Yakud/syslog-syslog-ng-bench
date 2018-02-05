NAME:=logging-system
MAINTAINER:="Aleksei Kiselev <yakud@kit.ec>"
DESCRIPTION:="logging-system"

all: run

init:
	docker network create logs-network && cd ./log-generator-image && sh ./build.sh

run:
	docker-compose -f compose.yml up --force-recreate

run-daemon:
	docker-compose -f compose.yml up -d --force-recreate