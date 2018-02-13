FROM debian

COPY gen-log /usr/local/bin

ENTRYPOINT ["/usr/local/bin/gen-log", "-syslog=0", "-stdout=1", "-workers=48"]