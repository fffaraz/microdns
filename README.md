# MicroDNS
A tiny DNS server in Go that always returns the same IP for any query sent to it.

cat log.out | cut -f1 | sort | uniq -cd | sort -nr
