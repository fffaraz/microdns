# MicroDNS
A tiny dockerized DNS server in Go that always returns the same IP for any query sent to it.

cat log.out | cut -f1 | sort | uniq -cd | sort -nr

docker build -t fffaraz/microdns:latest -f Dockerfile .

docker run -d --restart=always --name microdns -p 53:53 -p 53:53/udp fffaraz/microdns -ipv4 127.0.0.1 -ipv6 ::1 -ttl 86400
