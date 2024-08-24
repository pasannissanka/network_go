package client_test

import (
	"fmt"
	"testing"

	"github.com/pasannissanka/network_go/lib/client"
)

func TestIP(t *testing.T) {
	t.Log("Testing IP to CIDR conversion")

	cidr, err := client.GetCIDRs("192.168.1.18/32")

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("CIDR: %v\n", cidr)

	t.Log("CIDR: ", cidr)

	// cidrs := client.CalculateSubnetIPs("192.168.1.18/32", 10)

	// t.Logf("CIDRs: %s", cidrs)
}
