package client

import (
	"fmt"
	"log"

	"net"
	"sync"
	"time"

	"github.com/pasannissanka/network_go/lib/server"
)

var (
	CIDRs     []string
	port      uint16 = server.ServerData.UDP_L
	portStart uint16 = 1
	portEnd   uint16 = 65535
	isTest    bool   = true
	wg        sync.WaitGroup
	timeout   = time.Microsecond * 2000
)

type TestModeOptions struct {
	PortStart uint16
	PortEnd   uint16
}

func EnableTestMode(options *TestModeOptions) {
	log.Println("Test mode enabled")

	isTest = true

	if options.PortStart == 0 {
		options.PortStart = 1
	}
	if options.PortEnd == 0 {
		options.PortEnd = 65535
	}

	portStart = options.PortStart
	portEnd = options.PortEnd
}

func Scan(ip string) {
	log.Println("Scanning...")
	// If the test flag is set, use the server IP to scan
	// CIDRs, err := GetCIDRs(ip)

	// if err != nil {
	// 	fmt.Printf("not able to parse 'ips' parameter value: %s\n", err)
	// }

	if !isTest {
		portStart = port
		portEnd = port
	}

	// fmt.Printf("cidrs: %v\n", CIDRs)
	// for _, cidr := range CIDRs {
	// 	fmt.Printf("Scanning CIDR: %s\n", cidr)
	// }

	scan(ip)
}

func scan(cidr string) (err error) {
	var ip net.IP
	var ipNet *net.IPNet

	var incIP = func(ip net.IP) {
		for j := len(ip) - 1; j >= 0; j-- {
			ip[j]++
			if ip[j] > 0 {
				break
			}
		}
	}

	ip, ipNet, err = net.ParseCIDR(cidr)

	log.Printf("ip: %v\n", ip)
	log.Printf("ipNet: %v\n", ipNet)

	if err != nil {
		log.Printf("CIDR address not in correct format: %s", err)
		return err
	}

	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()

			for port := portStart; port <= portEnd; port++ {
				addr := fmt.Sprintf("%s:%d", ip, port)
				log.Printf("scanning addr: %s://%s\n", "udp", addr)

				c, e := net.DialTimeout("udp", addr, timeout)
				if e == nil {
					log.Printf("udp://%s is alive and reachable\n", addr)
					err := Connect(c)

					if err == nil {
						log.Printf("Connection to %s successful\n", addr)
					} else {
						log.Printf("Connection to %s failed: %s\n", addr, err)
					}
				}

			}
		}(ip.String())
	}

	wg.Wait()

	return err
}
