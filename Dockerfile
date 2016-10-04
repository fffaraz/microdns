FROM golang:alpine
MAINTAINER Faraz Fallahi <fffaraz@gmail.com>
RUN apk add --update git && rm -rf /var/cache/apk/* && go get github.com/miekg/dns
RUN git clone https://github.com/fffaraz/microdns && cd microdns && git checkout v2 && go build
EXPOSE 53/udp 53/tcp
ENTRYPOINT ["/microdns/microdns"]
