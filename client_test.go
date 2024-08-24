package main_test

import (
	"testing"

	"fmt"
	"net"
	"time"

	"github.com/pasannissanka/network_go/client"
)

func TestClientConnect(t *testing.T) {
	timeout := time.Microsecond * 2000
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 8856)

	c, e := net.DialTimeout("udp", addr, timeout)

	if e != nil {
		t.Error(e)
	}

	if c == nil {
		t.Error("Connection failed")
	}

	fmt.Printf("Connection to %s successful\n", addr)

	client.Connect(c)
}
