package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func handleConnection(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())

	conn.Write([]byte("Welcome to the server!\n"))
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		fmt.Println(temp)

		result := strconv.Itoa(10) + "\n"
		conn.Write([]byte(string(result)))
	}

	defer conn.Close()
}

func tcp() {
	fmt.Printf("Starting TCP server on ip: %s , port : %d \n", ServerData.IP, ServerData.TCP)

	ln, err := net.Listen("tcp", fmt.Sprintf(": %d", ServerData.TCP))

	if err != nil {
		fmt.Printf("Error starting TCP server: %s \n", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}

		handleConnection(conn)
	}
}
