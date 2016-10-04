FROM golang:alpine
MAINTAINER Faraz Fallahi <fffaraz@gmail.com>
EXPOSE 53/udp 53/tcp
ENTRYPOINT ["/go/microdns"]
RUN apk add --update git && rm -rf /var/cache/apk/* && go get github.com/miekg/dns && git clone -b v2 --depth 1 https://github.com/fffaraz/microdns /go/src/microdns && go build microdns
