package dnsbl

import (
	"math/rand"
	"time"
)

var nameservers = []string{"8.8.8.8:53", "8.8.4.4:53", "1.1.1.1:53", "208.67.222.222:53"}

type DNSBLConfig struct {
	Nameservers []string
	Concurrency int
	Dnsbl       string
	Debug       bool
}

func GetDefaultConfig() DNSBLConfig {
	return DNSBLConfig{
		Nameservers: nameservers,
		Concurrency: 5,
		Debug:       false,
	}
}

func (c DNSBLConfig) nameserver() string {
	if len(c.Nameservers) > 1 {
		rand.Seed(time.Now().Unix())
		return c.Nameservers[rand.Intn(len(c.Nameservers))]
	}
	return c.Nameservers[0]
}

func (c DNSBLConfig) RandomNameserver() string {
	return c.nameserver()
}

var dnsbls = []string{
	"black.uribl.com",
	"dnsbl-1.uceprotect.net",
	"dnsbl-2.uceprotect.net",
	"dnsbl-3.uceprotect.net",
	"gray.uribl.com",
	"mail.bl.blocklist.de",
	"access.redhawk.org",
	"b.barracudacentral.org",
	"bl.0spam.org ",
	"bl.spamcop.net",
	"blackholes.mail-abuse.org",
	"bogons.cymru.com",
	"cbl.abuseat.org",
	"cdl.anti-spam.org.cn",
	"ivmuri.ivmsip.com",
	"ivmuri.ivmuri.com",
	"ivmuri.ivmsip.org",
	"csi.cloudmark.com",
	"db.wpbl.info",
	"dnsbl.dronebl.org",
	"dnsbl.inps.de",
	"dnsbl.njabl.org",
	"dnsbl.sorbs.net",
	"drone.abuse.ch",
	"dsn.rfc-ignorant.org",
	"dul.dnsbl.sorbs.net",
	"dyna.spamrats.com",
	"httpbl.abuse.ch",
	"ips.backscatterer.org",
	"ix.dnsbl.manitu.net",
	"korea.services.net",
	"misc.dnsbl.sorbs.net",
	"multi.surbl.org",
	"netblock.pedantic.org",
	"noptr.spamrats.com",
	"opm.tornevall.org",
	"pbl.spamhaus.org",
	"dbl.spamhaus.org",
	"psbl.surriel.com",
	"query.senderbase.org",
	"rbl-plus.mail-abuse.org",
	"rbl.efnetrbl.org",
	"rbl.interserver.net",
	"rbl.spamlab.com",
	"rbl.suresupport.com",
	"relays.mail-abuse.org",
	"sbl.spamhaus.org",
	"short.rbl.jp",
	"smtp.dnsbl.sorbs.net",
	"socks.dnsbl.sorbs.net",
	"spam.dnsbl.sorbs.net",
	"spam.spamrats.com",
	"spamguard.leadmon.net",
	"spamrbl.imp.ch",
	"tor.dan.me.uk",
	"ubl.unsubscore.com",
	"virbl.bit.nl",
	"virus.rbl.jp",
	"web.dnsbl.sorbs.net",
	"wormrbl.imp.ch",
	"xbl.abuseat.org",
	"xbl.spamhaus.org",
	"zen.spamhaus.org",
	"zombie.dnsbl.sorbs.net",
}

func getDNSBLs() []string {
	return dnsbls
}
