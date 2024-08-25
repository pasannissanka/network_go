package main

import (
	"fmt"
	"time"

	"github.com/pasannissanka/network_go/lib/client"
)

func ConnectToMaster() {
	client.EnableTestMode(&client.TestModeOptions{
		PortStart: 8881,
		PortEnd:   8881,
	})

	for {
		client.Scan(fmt.Sprintf("%s/24", Env.Ip))
		time.Sleep(1 * time.Minute)
	}
}

func PublishMessages() {
	connections := client.TcpConnections.GetConnections()

	for _, conn := range connections {
		go func(conn client.Connection) {

		}(conn)
	}
}
