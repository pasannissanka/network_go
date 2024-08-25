package main

import (
	"fmt"
	"time"

	"github.com/pasannissanka/network_go/lib/client"
	server_net "github.com/pasannissanka/network_go/lib/net"
)

func ConnectToMaster() {
	client.EnableTestMode(&client.TestModeOptions{
		PortStart: 8880,
		PortEnd:   8885,
	})

	for {
		client.Scan(fmt.Sprintf("%s/24", Env.Ip), Env.Id)
		time.Sleep(1 * time.Minute)
	}
}

func PublishMessages() {
	connections := client.TcpConnections.GetConnections()

	for _, conn := range connections {
		go func(conn server_net.Connection) {

		}(conn)
	}
}
