# DNSBL Blacklist Checker

*Some of the code is borrowed from [BlackList Checker by ilijamt](https://github.com/ilijamt/blacklist-checker)*

Check if an IP v4 is blacklisted. 

![Screenshot of My Project](./images/dnsbl-console.png)

## Requirements

- GoLang >= 1.20

## Installation

```console
foo@bar:~$ git clone https://github.com/igorrendulic/go-dnsbl-blacklist-checker.git
foo@bar:~$ cd go-dnsbl-blacklist-checker
foo@bar:~$ make build
```

## Usage

```console
foo@bar:~$ ./dnsbl-blacklist-checker 127.0.0.1
```

Use as module in your own project:

```
import "github.com/mailio/go-dnsbl-blacklist-checker"
```

Simple use:
```go
// get config
cfg := dnsbl.GetDefaultConfig()
nameserver := cfg.RandomNameserver()

// run DNSBL queries
results, err := dnsbl.RunQuery(cfg, ip, nameserver)
```

## Blacklists

find them in `dnsbl/config.go`

```
black.uribl.com
dnsbl-1.uceprotect.net
dnsbl-2.uceprotect.net
dnsbl-3.uceprotect.net
gray.uribl.com
mail.bl.blocklist.de
access.redhawk.org
b.barracudacentral.org
bl.0spam.org 
bl.spamcop.net
blackholes.mail-abuse.org
bogons.cymru.com
cbl.abuseat.org
cdl.anti-spam.org.cn
ivmuri.ivmsip.com
ivmuri.ivmuri.com
ivmuri.ivmsip.org
csi.cloudmark.com
db.wpbl.info
dnsbl.dronebl.org
dnsbl.inps.de
dnsbl.njabl.org
dnsbl.sorbs.net
drone.abuse.ch
dsn.rfc-ignorant.org
dul.dnsbl.sorbs.net
dyna.spamrats.com
httpbl.abuse.ch
ips.backscatterer.org
ix.dnsbl.manitu.net
korea.services.net
misc.dnsbl.sorbs.net
multi.surbl.org
netblock.pedantic.org
noptr.spamrats.com
opm.tornevall.org
pbl.spamhaus.org
dbl.spamhaus.org
psbl.surriel.com
query.senderbase.org
rbl-plus.mail-abuse.org
rbl.efnetrbl.org
rbl.interserver.net
rbl.spamlab.com
rbl.suresupport.com
relays.mail-abuse.org
sbl.spamhaus.org
short.rbl.jp
smtp.dnsbl.sorbs.net
socks.dnsbl.sorbs.net
spam.dnsbl.sorbs.net
spam.spamrats.com
spamguard.leadmon.net
spamrbl.imp.ch
tor.dan.me.uk
ubl.unsubscore.com
virbl.bit.nl
virus.rbl.jp
web.dnsbl.sorbs.net
wormrbl.imp.ch
xbl.abuseat.org
xbl.spamhaus.org
zen.spamhaus.org
zombie.dnsbl.sorbs.net
```


## More DNSBL servers if needed

- [blacklist-check-unix-linux-utility by adionditsak](https://github.com/adionditsak/blacklist-check-unix-linux-utility)