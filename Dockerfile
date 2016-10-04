FROM golang:alpine
MAINTAINER Faraz Fallahi <fffaraz@gmail.com>
RUN apk add --update git && rm -rf /var/cache/apk/* 
RUN git clone github.com/fffaraz/microdns && cd microdns && git checkout v2 && go build
EXPOSE 53/udp 53/tcp
ENTRYPOINT ["/microdns/microdns"]
