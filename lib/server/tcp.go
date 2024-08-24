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
	log.Printf("Serving %s\n", conn.RemoteAddr().String())

	conn.Write([]byte("Welcome to the server!\n"))
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Fatal(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		log.Println(temp)

		result := strconv.Itoa(10) + "\n"
		conn.Write([]byte(string(result)))
	}

	defer conn.Close()
}

func tcp() {
	log.Printf("Starting TCP server on ip: %s , port : %d \n", ServerData.IP, ServerData.TCP)

	ln, err := net.Listen("tcp", fmt.Sprintf(": %d", ServerData.TCP))

	if err != nil {
		log.Panicf("Error starting TCP server: %s \n", err)
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		handleConnection(conn)
	}
}
