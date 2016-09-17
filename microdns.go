package main

import (
	"flag"
	//"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"github.com/miekg/dns"
)

var ipv4, ipv6 string
var ttl int

func main() {
	flag.StringVar(&ipv4, "ipv4", "127.0.0.1", "IPv4 Address")
	flag.StringVar(&ipv6, "ipv6", "::1", "IPv6 Address")
	flag.IntVar(&ttl, "ttl", 86400, "Time to live")
	flag.Parse()
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
	//ip, _, _ := net.SplitHostPort(w.RemoteAddr().String())
	//fmt.Printf("%s\t%s\n", ip, r.Question[0].Name)
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true
	rr1 := new(dns.A)
	rr1.Hdr = dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: uint32(ttl)}
	rr1.A = net.ParseIP(ipv4)
	rr2 := new(dns.AAAA)
	rr2.Hdr = dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: uint32(ttl)}
	rr2.AAAA = net.ParseIP(ipv6)
	m.Answer = []dns.RR{rr1, rr2}
	w.WriteMsg(m)
}
