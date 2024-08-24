package server

import (
	"fmt"
	"log"
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
	TCP   uint16 // TCP port
}

var TCP_PORT Port
var UDP_L_PORT Port
var ServerData Server

var ID int

func Init(serverData Server, id int) {
	ServerData = serverData
	// Initialize the server data
	TCP_PORT = Port{
		IP:   serverData.IP,
		PORT: fmt.Sprint(serverData.TCP),
		DATA: "TCP",
	}

	UDP_L_PORT = Port{
		IP:   serverData.IP,
		PORT: fmt.Sprint(serverData.UDP_L),
		DATA: "UDP_L",
	}

	ID = id

	// Start the TCP and UDP servers
	defer tcp()
	defer udp()

	log.Println("Server started, press <CTRL+C> to exit")
}
