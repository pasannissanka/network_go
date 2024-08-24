package client

import (
	"fmt"

	"net"
	"sync"
	"time"

	"github.com/pasannissanka/network_go/server"
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

func EnableTestMode() {
	isTest = true
	portStart = 8850
	portEnd = 8860
}

func Scan(ip string) {
	fmt.Println("Scanning...")
	// If the test flag is set, use the server IP to scan
	if isTest {
		ip = "127.0.0.1/32"
	}

	CIDRs, err := GetCIDRs(ip)

	if err != nil {
		fmt.Printf("not able to parse 'ips' parameter value: %s\n", err)
	}

	if !isTest {
		portStart = port
		portEnd = port
	}

	fmt.Printf("cidrs: %v\n", CIDRs)
	for _, cidr := range CIDRs {
		fmt.Printf("Scanning CIDR: %s\n", cidr)
		scan(cidr)
	}

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

	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ipNet: %v\n", ipNet)

	if err != nil {
		fmt.Printf("CIDR address not in correct format: %s", err)
		return err
	}

	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()

			for port := portStart; port <= portEnd; port++ {
				addr := fmt.Sprintf("%s:%d", ip, port)
				fmt.Printf("scanning addr: %s://%s\n", "udp", addr)

				c, e := net.DialTimeout("udp", addr, timeout)
				if e == nil {
					fmt.Printf("udp://%s is alive and reachable\n", addr)
					err := Connect(c)

					if err == nil {
						fmt.Printf("Connection to %s successful\n", addr)
					} else {
						fmt.Printf("Connection to %s failed: %s\n", addr, err)
					}
				}

			}
		}(ip.String())
	}

	wg.Wait()

	return err
}
