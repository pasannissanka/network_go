package server

import (
	"fmt"
)

type Port struct {
	IP   string
	PORT string
	DATA string
}

type Message struct {
	IP   string
	PORT string
	ID   int
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

var TCP_PORT Port
var UDP_L_PORT Port

var ID int

func Init(id int) {
	// Initialize the server data
	TCP_PORT = Port{
		IP:   ServerData.IP,
		PORT: fmt.Sprint(ServerData.TCP),
		DATA: "TCP",
	}

	UDP_L_PORT = Port{
		IP:   ServerData.IP,
		PORT: fmt.Sprint(ServerData.UDP_L),
		DATA: "UDP_L",
	}

	ID = id

	// Start the TCP and UDP servers
	defer tcp()
	go udp()
}
