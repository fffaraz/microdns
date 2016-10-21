FROM golang:alpine
MAINTAINER Faraz Fallahi <fffaraz@gmail.com>
EXPOSE 53/udp 53/tcp
ENTRYPOINT ["microdns"]
RUN \
	apk add --update --no-cache git && \
	go get github.com/fffaraz/microdns && \
	apk del libssh2 libcurl expat pcre git && \
	rm -rf /go/pkg /go/src /var/cache/apk
