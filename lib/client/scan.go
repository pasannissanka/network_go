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

func Scan(ip string) {
	// If the test flag is set, use the server IP to scan
	if isTest {
		ip = server.ServerData.IP
	}

	CIDRs, err := GetCIDRs(ip)

	if err != nil {
		fmt.Printf("not able to parse 'ips' parameter value: %s\n", err)
	}

	if !isTest {
		portStart = port
		portEnd = port
	}

	for _, cidr := range CIDRs {
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

	if err != nil {
		fmt.Printf("CIDR address not in correct format %s", err)
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
					Connect(c)
				}

			}
		}(ip.String())
	}

	wg.Wait()

	return err
}
