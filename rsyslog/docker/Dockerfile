FROM alpine:3.6

COPY	rsyslog@lists.adiscon.com-5a55e598.rsa.pub /etc/apk/keys/rsyslog@lists.adiscon.com-5a55e598.rsa.pub
RUN	echo "http://build.rsyslog.com/alpine/3.7/unstable" >> /etc/apk/repositories \
	&& apk --no-cache update  \
	&& apk add --no-cache \
	   rsyslog \
	   rsyslog-elasticsearch \
	   rsyslog-imptcp \
	   rsyslog-imrelp \
	   rsyslog-omrelp \
	   rsyslog-omstdout \
	   rsyslog-omhttpfs


COPY . /
VOLUME /var/run/rsyslog/dev
ENTRYPOINT ["rsyslogd", "-n", "-f", "/etc/rsyslogd.conf"]

EXPOSE 514/tcp
EXPOSE 514/udp