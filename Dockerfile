FROM golang:alpine AS builder
RUN apk add --update git && go get github.com/fffaraz/microdns

FROM alpine:latest
COPY --from=builder /go/bin/microdns /usr/local/bin
ENTRYPOINT ["microdns"]
