package main

import (
	"github.com/miekg/dns"

	"encoding/json"
	"strings"
	"time"
)

type RRecord struct {
	Name string `json:"name"`
	Ttl  uint32 `json:"ttl"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func GetZone(zone string, server string, tsig Tsig) ([]byte, error) {
	t := new(dns.Transfer)
	m := new(dns.Msg)

	t.TsigSecret = tsig.Map

	m.SetAxfr(zone)
	m.SetTsig(tsig.Name, dns.HmacSHA256, 300, time.Now().Unix())
	c, err := t.In(m, server)

	if err != nil {
		return nil, err
	}
	value := <-c

	var response []RRecord

	for _, record := range value.RR {
		if !(record.Header().Rrtype == dns.TypeSOA) {
			split := strings.Fields(record.String())

			response = append(response, RRecord{
				record.Header().Name,
				record.Header().Ttl,
				dns.Type.String(dns.Type(record.Header().Rrtype)),
				split[len(split)-1]})
		}
	}
	jsonrep, err := json.Marshal(response)
	return jsonrep, err
}
