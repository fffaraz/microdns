FROM golang:alpine
MAINTAINER Faraz Fallahi <fffaraz@gmail.com>
EXPOSE 53/udp 53/tcp
ENTRYPOINT ["/go/microdns/microdns"]
RUN apk add --update git && rm -rf /var/cache/apk/* && go get github.com/miekg/dns && git clone -b v2 https://github.com/fffaraz/microdns && go build microdns
