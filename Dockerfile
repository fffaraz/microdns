FROM golang:alpine
MAINTAINER Faraz Fallahi <fffaraz@gmail.com>
RUN apk add --update git && rm -rf /var/cache/apk/* && go get github.com/fffaraz/microdns
EXPOSE 53/udp 53/tcp
ENTRYPOINT ["microdns"]
