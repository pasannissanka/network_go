package main

import (
	"fmt"

	"github.com/pasannissanka/network_go/server"
)

func main() {
	fmt.Println("Hello, World!")

	serverData := server.Server{
		IP:    "127.0.0.1",
		UDP_L: 8856,
		TCP:   8881,
	}

	server.Init(serverData, 1)
}
