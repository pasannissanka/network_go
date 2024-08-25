package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func handleConnection(conn net.Conn) {
	addr := conn.RemoteAddr()
	log.Printf("[TCP] Serving %s\n", addr.String())

	conn.Write([]byte("Welcome to the server!\n"))
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Printf("[TCP][%s][ERROR]>: %s\n", addr.String(), err)
			break
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			log.Printf("[TCP][%s] Exiting TCP server!", addr.String())
			break
		}

		log.Printf("[TCP][%s]> %s", addr.String(), temp)

		result := strconv.Itoa(10) + "\n"
		conn.Write([]byte(string(result)))
	}

	defer conn.Close()
}

func tcp() {
	log.Printf("[TCP] Starting TCP server on ip: %s , port : %d \n", ServerData.IP, ServerData.TCP)

	ln, err := net.Listen("tcp", fmt.Sprintf(": %d", ServerData.TCP))

	if err != nil {
		log.Panicf("[TCP] Error starting TCP server: %s \n", err)
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
