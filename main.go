package main

import (
	"fmt"

	"net"

	"github.com/pasannissanka/network_go/lib/client"
)

func main() {
	// timeout := time.Microsecond * 2000
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 8856)

	fmt.Printf("Address to connect: %s\n", addr)

	c, e := net.Dial("udp", addr)

	if e != nil {
		fmt.Print(e)
	}

	if c == nil {
		fmt.Errorf("Connection failed")
	}

	fmt.Printf("Connection to %s successful\n", addr)

	client.Connect(c)
}
