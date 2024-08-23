package net_test

import (
	"fmt"
	"testing"

	"github.com/pasannissanka/network_go/net"
)

func TestIP(t *testing.T) {
	t.Log("Testing IP to CIDR conversion")

	cidr, err := net.GetCIDRs("127.0.0.1")

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("CIDR: %v\n", cidr)

	t.Log("CIDR: ", cidr)
}
