package models

import "net"

type DNSBL struct {
	IP                 string   `json:"ip"`
	Nameserver         string   `json:"nameserver"`           // which nameserver was used for query
	ARecordResponses   []string `json:"a_record_responses"`   // what was the response from the nameserver for the A record query
	TXTRecordResponses []string `json:"txt_record_responses"` // what was the response from the nameserver for the TXT record query
	Blacklisted        bool     `json:"blacklisted"`          // was the IP blacklisted
	Host               string   `json:"host"`                 // what blacklist was used
	Error              string   `json:"string"`               // error if any
	QueryFailed        bool     `json:"query_failed"`         // did the query fail
}

type Item struct {
	IP        net.IP
	Blacklist string
	Host      string
}
