package main

import (
	"fmt"
	"time"

	"github.com/miekg/dns"
)

func PostRecord(zone string, server string, tsig Tsig, name string, ttl uint32, rtype string, data string) {
	c := new(dns.Client)
	c.TsigSecret = tsig.Map
	m := new(dns.Msg)
	m.SetUpdate(zone)
	m.SetTsig(tsig.Name, dns.HmacSHA256, 300, time.Now().Unix())
	var rrset []dns.RR
	rr, _ := dns.NewRR(fmt.Sprintf("%s %d IN %s %s", name, ttl, rtype, data))
	rrset = append(rrset, rr)
	m.Insert(rrset)
	message, _, err := c.Exchange(m, server)

	fmt.Println("message: ", message)
	fmt.Println("error: ", err)
}
