package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/mailio/go-dnsbl-blacklist-checker/dnsbl"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <IPv4 address>\n", os.Args[0])
		os.Exit(1)
	}

	ipInput := os.Args[1]
	if net.ParseIP(ipInput) == nil || net.ParseIP(ipInput).To4() == nil {
		fmt.Fprintf(os.Stderr, "Invalid IPv4 address: %s\n", ipInput)
		os.Exit(1)
	}
	ip := net.ParseIP(ipInput)

	cfg := dnsbl.GetDefaultConfig()
	nameserver := cfg.RandomNameserver()

	results, err := dnsbl.RunQuery(cfg, ip, nameserver)
	if err != nil {
		panic(err)
	}
	for _, res := range results {
		aResp := strings.Join(res.ARecordResponses, ";")
		txtResp := strings.Join(res.TXTRecordResponses, ";")
		if res.Blacklisted {
			fmt.Printf("\033[31m%s is blacklisted in %s, A: %s, TXT: %s using NS: %s\033[0m\n", ip.String(), res.Host, aResp, txtResp, res.Nameserver)
		} else if res.QueryFailed {
			fmt.Printf("\033[33m%s query failed in %s, A: %s, TXT: %s using NS: %s\033[0m\n", ip.String(), res.Host, aResp, txtResp, res.Nameserver)
		} else {
			fmt.Printf("%s is not blacklisted in %s, A: %s, TXT: %s using NS: %s\n", ip.String(), res.Host, aResp, txtResp, res.Nameserver)
		}
	}
}
