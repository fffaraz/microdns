package main

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
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
	fmt.Printf("%s,%s\n", w.RemoteAddr().String(), r.Question[0].Name)
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true
	rr1 := new(dns.A)
	rr1.Hdr = dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 86400}
	rr1.A = net.ParseIP("127.0.0.1")
	rr2 := new(dns.AAAA)
	rr2.Hdr = dns.RR_Header{Name: r.Question[0].Name, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 86400}
	rr2.AAAA = net.ParseIP("::1")
	m.Answer = []dns.RR{rr1, rr2}
	w.WriteMsg(m)
}
