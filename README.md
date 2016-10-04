# MicroDNS
A tiny dockerized DNS server in Go that always returns the same IP for any query sent to it.

```
docker run -d --restart=always -p 53:53 -p 53:53/udp --name microdns \
fffaraz/microdns:latest \
-ipv4 127.0.0.1 \
-ipv6 ::1 \
-ttl 86400 \
-log true
```

```
docker run -d --restart=always -p 53:53 -p 53:53/udp --name microdns \
-v /home/microdns:/home \
fffaraz/microdns:latest \
-ipv4 127.0.0.1 \
-ipv6 ::1 \
-ttl 86400 \
-log true \
-conf /home/dns.conf
```

```
cat log.out | cut -f1 | sort | uniq -cd | sort -nr
```
