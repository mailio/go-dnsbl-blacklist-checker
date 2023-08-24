package dnsbl

import (
	"context"
	"strings"
	"time"

	"github.com/mailio/mailio-dofindserver/models"
	"github.com/miekg/dns"
)

// check the description (if may not exist for the domain)
func queryTXT(item models.Item, nameserver string) (string, error) {
	client := new(dns.Client)
	var t = dns.TypeTXT
	m := new(dns.Msg)

	m.SetQuestion(dns.Fqdn(item.Blacklist), t)
	m.RecursionAvailable = true

	var r *dns.Msg
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, _, err := client.ExchangeContext(ctx, m, nameserver)
	if err != nil {
		return "", err
	}
	if len(r.Answer) <= 0 {
		return "", nil
	}
	response := ""
	for _, a := range r.Answer {
		switch a.(type) {
		case *dns.TXT:
			txtAnswer := a.(*dns.TXT).Txt
			response = strings.Join(txtAnswer, "; ")
		}
	}
	return response, nil
}

// query A records in the DNSBL
func checkDNSBL(item models.Item, nameserver string) (blacklisted bool, aResponses []string, txtResponses []string, err error) {

	client := new(dns.Client)
	isIPv6 := item.IP.To4() == nil

	var t = dns.TypeA
	m := new(dns.Msg)
	if isIPv6 {
		t = dns.TypeAAAA
	}
	m.SetQuestion(dns.Fqdn(item.Blacklist), t)
	m.RecursionAvailable = true

	var r *dns.Msg
	r, _, err = client.Exchange(m, nameserver)
	if err != nil {
		return blacklisted, aResponses, txtResponses, err
	}

	if len(r.Answer) > 0 {
		blacklisted = true
		txtResp, txtErr := queryTXT(item, nameserver)
		if txtErr != nil {
			return blacklisted, aResponses, txtResponses, txtErr
		}
		txtResponses = append(txtResponses, txtResp)
	}

	for _, a := range r.Answer {
		switch a.(type) {
		case *dns.A:

			aResponses = append(aResponses, a.(*dns.A).A.String())
		case *dns.AAAA:
			aResponses = append(aResponses, a.(*dns.AAAA).AAAA.String())
		}
	}

	return blacklisted, aResponses, txtResponses, nil

}
