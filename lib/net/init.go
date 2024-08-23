package net

import (
	"fmt"
)

type Message struct {
	IP   string
	PORT string
	DATA string
}

type Server struct {
	IP    string // IP address of the server
	UDP_L uint16 // UDP listener port
	UDP_R uint16 // UDP receiver port
	TCP   uint16 // TCP port
}

var ServerData = Server{
	IP:    "127.0.0.1",
	UDP_L: 8888,
	UDP_R: 8889,
	TCP:   8881,
}

var TCP_PORT Message
var UDP_L_PORT Message
var UDP_R_PORT Message

func Init() {
	// Initialize the server data
	TCP_PORT = Message{
		IP:   ServerData.IP,
		PORT: fmt.Sprint(ServerData.TCP),
		DATA: "TCP",
	}

	UDP_L_PORT = Message{
		IP:   ServerData.IP,
		PORT: fmt.Sprint(ServerData.UDP_L),
		DATA: "UDP_L",
	}

	UDP_R_PORT = Message{
		IP:   ServerData.IP,
		PORT: fmt.Sprint(ServerData.UDP_R),
		DATA: "UDP_R",
	}

	// Start the TCP and UDP servers
	defer tcp()
}
