FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" .

FROM scratch
COPY --from=builder /app/microdns /usr/local/bin
EXPOSE 53
ENTRYPOINT ["microdns"]
