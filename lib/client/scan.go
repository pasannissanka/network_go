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
	ID        int
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

func Scan(ip string, id int) {
	log.Println("Scanning...")

	ID = id

	log.Printf("ID: %d\n", ID)

	if !isTest {
		portStart = port
		portEnd = port
	}

	scan(ip)
}

func scan(cidr string) (err error) {
	log.Printf("Scanning CIDR: %s\n", cidr)

	var ip net.IP
	var ipNet *net.IPNet

	ip, ipNet, err = net.ParseCIDR(cidr)

	log.Printf("ip: %v\n", ip)
	log.Printf("ipNet: %v\n", ipNet)

	if err != nil {
		log.Printf("CIDR address not in correct format: %s", err)
		return err
	}

	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()

			for port := portStart; port <= portEnd; port++ {
				addr := fmt.Sprintf("%s:%d", ip, port)
				c, e := net.DialTimeout("udp", addr, timeout)
				if e == nil {
					err := Connect(c)

					if err == nil {
						log.Printf("Connection to %s successful\n", addr)
					}
				}

			}
		}(ip.String())
	}
	wg.Wait()

	log.Print("Scan complete\n")
	return err
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
