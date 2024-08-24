package server

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	bin_buf := new(bytes.Buffer)

	data := Message{
		IP:   ServerData.IP,
		PORT: fmt.Sprint(ServerData.TCP),
		ID:   ID,
	}
	// create a encoder object
	gobobj := gob.NewEncoder(bin_buf)
	// encode buffer and marshal it into a gob object
	gobobj.Encode(data)

	_, err := conn.WriteToUDP(bin_buf.Bytes(), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func udp() {
	fmt.Printf("Starting UDP server on ip: %s port : %d \n", ServerData.IP, ServerData.UDP_L)

	addr := net.UDPAddr{
		Port: int(ServerData.UDP_L),
		IP:   net.ParseIP(ServerData.IP),
	}

	conn, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	p := make([]byte, 2048)

	for {
		_, remote_addr, err := conn.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remote_addr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}

		sendResponse(conn, remote_addr)
	}
}
