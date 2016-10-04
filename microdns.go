package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"github.com/miekg/dns"
)

var ipv4, ipv6, conf string
var ttl int
var logflag bool
var mapv4 map[string]string
var mapv6 map[string]string

func main() {
	flag.StringVar(&ipv4, "ipv4", "127.0.0.1", "IPv4 Address")
	flag.StringVar(&ipv6, "ipv6", "::1", "IPv6 Address")
	flag.IntVar(&ttl, "ttl", 86400, "Time to live")
	flag.BoolVar(&logflag, "log", false, "Log to stdout")
	flag.StringVar(&conf, "conf", "", "Config File")
	flag.Parse()
	file, err := os.Open(conf)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		mapv4 = make(map[string]string)
		mapv6 = make(map[string]string)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			fields := strings.Split(line, "\t")
			mapv4[fields[0]] = fields[1]
			mapv6[fields[0]] = fields[2]
		}
	}
	fmt.Println("")
	dns.HandleFunc(".", handleRequest)
	go func() {
		srv := &dns.Server{Addr: ":53", Net: "udp"}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to set udp listener %s\n", err.Error())
		}
	}()
	go func() {
		srv := &dns.Server{Addr: ":53", Net: "tcp"}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to set tcp listener %s\n", err.Error())
		}
	}()
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-sig:
			log.Fatalf("Signal (%d) received, stopping\n", s)
		}
	}
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	domain := r.Question[0].Name
	if logflag {
		ip, _, _ := net.SplitHostPort(w.RemoteAddr().String())
		fmt.Printf("%s\t%s\n", ip, domain)
	}
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true
	rr1 := new(dns.A)
	rr1.Hdr = dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: uint32(ttl)}
	rr2 := new(dns.AAAA)
	rr2.Hdr = dns.RR_Header{Name: domain, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: uint32(ttl)}
	if val, ok := mapv4[domain]; ok {
		rr1.A = net.ParseIP(val)
		rr2.AAAA = net.ParseIP(mapv6[domain])
	} else {
		rr1.A = net.ParseIP(ipv4)
		rr2.AAAA = net.ParseIP(ipv6)
	}
	m.Answer = []dns.RR{rr1, rr2}
	w.WriteMsg(m)
}
