# MicroDNS
A tiny dockerized DNS server in Go that always returns the same IP for any query sent to it.

## Options

```
	-ipv4		IPv4 to return (default: 127.0.0.1)
	-ipv6		IPv6 to return (default: ::1)
	-ttl		Time to live (default: 86400)
	-log		Log requests to stdout (default: false)
	-conf		Config file (default: /home/dns.conf)
```

## How to Run

```
docker run -d --restart=always -p 53:53 -p 53:53/udp --name microdns \
fffaraz/microdns:latest \
-ipv4 127.0.0.1 \
-ipv6 ::1 \
-ttl 86400 \
-log
```

```
docker run -d --restart=always -p 53:53 -p 53:53/udp --name microdns \
-v /home/microdns:/home \
fffaraz/microdns:latest \
-ipv4 127.0.0.1 \
-ipv6 ::1 \
-ttl 86400 \
-log \
-conf /home/dns.conf
```

```
cat log.out | cut -f1 | sort | uniq -cd | sort -nr
```
