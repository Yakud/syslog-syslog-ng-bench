NAME:=logging-system
MAINTAINER:="Aleksei Kiselev <yakud@kit.ec>"
DESCRIPTION:="logging-system"

all: run

init:
	docker network create logs-network && cd ./log-generator-image && sh ./build.sh && \
	cd ../rsyslog/docker && sh build.sh

run-syslog-ng:
	docker-compose -f compose-syslog-ng.yml up --force-recreate

run-rsyslog:
	docker-compose -f compose-rsyslog.yml up --force-recreate

#run-daemon:
#	docker-compose -f compose.yml up -d --force-recreate