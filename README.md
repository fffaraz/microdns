# MicroDNS
A tiny dockerized DNS server in Go that (almost) always returns the same IP for any query sent to it.

[![](http://dockeri.co/image/fffaraz/microdns)](https://hub.docker.com/r/fffaraz/microdns/)
[![](https://images.microbadger.com/badges/image/fffaraz/microdns.svg)](https://microbadger.com/images/fffaraz/microdns)

## Quick Start (TL;DR)

```
docker run -it --rm -p 53:53 -p 53:53/udp fffaraz/microdns:latest -log \
-ipv4 $(dig +short myip.opendns.com @resolver1.opendns.com) \
-ipv6 $(dig +short myip.opendns.com @2620:0:ccc::2 aaaa)
```

## Arguments & Options

```
	-ipv4		Default IPv4 to return (default: 127.0.0.1)
	-ipv6		Default IPv6 to return (default: ::1)
	-ttl		Time to live value (default: 86400)
	-log		Log requests to stdout (default: false)
	-conf		Config file (default: /home/dns.conf)
```

## Config file format

```
domain1.com.     127.0.0.1 ::1
domain2.com.     127.0.0.1 ::1
www.domain1.com. 127.0.0.1 ::1
ftp.domain2.com. 127.0.0.1 ::1
```

## How to Run

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

## Request Stats

```
IP Adresses: cat log.out | cut -f2 | sort | uniq -cd | sort -nr
Domains:     cat log.out | cut -f3 | sort | uniq -cd | sort -nr
```

## Metadata API

```
curl http://169.254.169.254/metadata/v1/interfaces/public/0/ipv4/address
curl http://169.254.169.254/metadata/v1/interfaces/public/0/ipv6/address
```
