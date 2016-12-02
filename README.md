# MicroDNS
A tiny dockerized DNS server in Go that (almost) always returns the same IP for any query sent to it.

[![](http://dockeri.co/image/fffaraz/microdns)](https://hub.docker.com/r/fffaraz/microdns/)

## TL;DR

```
docker run -it --rm -p 53:53 -p 53:53/udp --name microdns \
fffaraz/microdns:latest \
-ipv4 $(dig +short myip.opendns.com @resolver1.opendns.com) \
-ipv6 $(dig +short myip.opendns.com @2620:0:ccc::2 aaaa) \
-log
```

## Options

```
	-ipv4		Default IPv4 to return (default: 127.0.0.1)
	-ipv6		Default IPv6 to return (default: ::1)
	-ttl		Time to live value (default: 86400)
	-log		Log requests to stdout (default: false)
	-conf		Config file (default: /home/dns.conf)
```

## Config file format

```
domain1.com. [tab] ipv4 [tab] ipv6
domain2.com. [tab] ipv4 [tab] ipv6
www.domain1.com. [tab] ipv4 [tab] ipv6
ftp.domain2.com. [tab] ipv4 [tab] ipv6
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
IP Adresses: cat log.out | cut -f1 | sort | uniq -cd | sort -nr
Domains:     cat log.out | cut -f2 | sort | uniq -cd | sort -nr
```
