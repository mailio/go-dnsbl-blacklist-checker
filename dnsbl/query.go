package dnsbl

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/mailio/mailio-dofindserver/models"
	"golang.org/x/sync/semaphore"
)

var (
	errorLogger *log.Logger
	debugLogger *log.Logger
)

func init() {
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func RunQuery(config DNSBLConfig, ip net.IP, nameserver string) ([]models.DNSBL, error) {
	hosts := getDNSBLs()

	if config.Debug {
		debugLogger.Printf("Querying %d blacklist hosts for %s\n", len(hosts), ip.String())
	}

	// modify concurrency to the length of hosts if necessary
	if config.Concurrency > len(hosts) {
		config.Concurrency = len(hosts)
	}
	if config.Debug {
		debugLogger.Printf("Concurrency set to %d\n", config.Concurrency)
	}

	sem := semaphore.NewWeighted(int64(config.Concurrency))
	var items []models.Item
	for _, host := range hosts {
		items = append(items, models.Item{
			IP:        ip,
			Blacklist: fmt.Sprintf("%s.%s.", reverseIP(ip.String()), host),
			Host:      host,
		})
		if config.Debug {
			debugLogger.Printf("Prepared query for %s\n", items[len(items)-1].Blacklist)
		}
	}

	var mu sync.Mutex
	dnsblResponses := make([]models.DNSBL, 0)

	// wait until all done
	var wg sync.WaitGroup

	for _, item := range items {
		if err := sem.Acquire(context.Background(), 1); err != nil {
			panic(err)
		}
		wg.Add(1)
		go func(i models.Item) {
			defer sem.Release(1) // relase the semaphore
			defer wg.Done()      // mark as done

			if config.Debug {
				debugLogger.Printf("Running %s\n", i.Blacklist)
			}

			isBlacklisted, aResponses, txtResponses, err := checkDNSBL(i, nameserver)
			if err != nil {
				if config.Debug {
					errorLogger.Printf("Error querying %s: %s\n", i.Blacklist, err.Error())
				}
				mu.Lock() // lock adding to the list
				defer mu.Unlock()
				dnsblResponses = append(dnsblResponses, models.DNSBL{
					IP:                 i.IP.String(),
					Blacklisted:        false,
					Nameserver:         nameserver,
					ARecordResponses:   aResponses,
					TXTRecordResponses: txtResponses,
					Host:               i.Host,
					Error:              err.Error(),
					QueryFailed:        true,
				})
				return
			}
			if config.Debug {
				debugLogger.Printf("Success querying %s returned %v\n", i.Blacklist, isBlacklisted)
			}
			mu.Lock() // lock adding to the list
			defer mu.Unlock()
			dnsblResponses = append(dnsblResponses, models.DNSBL{
				IP:                 i.IP.String(),
				Blacklisted:        isBlacklisted,
				Nameserver:         nameserver,
				ARecordResponses:   aResponses,
				TXTRecordResponses: txtResponses,
				Host:               i.Host,
				QueryFailed:        false,
			})
		}(item)
	}
	if config.Debug {
		debugLogger.Printf("Waiting for all queries to finish\n")
	}
	// Wait for all the semaphores to be released
	wg.Wait()
	// Process errors from error channel
	return dnsblResponses, nil
}
